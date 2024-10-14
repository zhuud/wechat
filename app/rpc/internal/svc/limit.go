package svc

import (
	"context"
	"fmt"
	"github.com/spf13/cast"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

const (
	burst       = 2
	minute      = 60
	hour        = 3600
	minuteQuota = 5000
	hourQuota   = 150000

	minuteKey = "wechat:token:second:limit:%s"
	hourKey   = "wechat:token:hour:limit:%s"
)

type (
	WechatTokenLimit struct {
		limiters map[string]*WechatLimiter
	}
	WechatLimiter [2]*limit.TokenLimiter
)

var (
	wl            *WechatTokenLimit
	wlOnce        sync.Once
	retryDuration = time.Minute
	keyQuota      = map[string][2]int{
		"all":           {minuteQuota, hourQuota},
		"external_user": {5000, 50000},
	}
)

func NewTokenWechatLimit(r *redis.Redis) *WechatTokenLimit {

	wlOnce.Do(func() {
		wl = &WechatTokenLimit{
			limiters: make(map[string]*WechatLimiter),
		}
		for k, q := range keyQuota {
			minuteBurst := q[0] / burst
			minuteRate := minuteBurst / minute

			hourBurst := q[1] / burst
			hourRate := hourBurst / hour

			ml := limit.NewTokenLimiter(minuteRate, minuteBurst, r, fmt.Sprintf(minuteKey, k))
			hl := limit.NewTokenLimiter(hourRate, hourBurst, r, fmt.Sprintf(hourKey, k))
			wl.limiters[k] = &WechatLimiter{ml, hl}
		}
	})

	return wl
}

func (wl *WechatTokenLimit) Allow(k string) bool {
	if k != "all" {
		l, _ := wl.limiters["all"]
		for _, limiter := range l {
			if !limiter.Allow() {
				return false
			}
		}
	}

	l, ok := wl.limiters[k]
	if !ok {
		return false
	}
	for _, limiter := range l {
		if !limiter.Allow() {
			return false
		}
	}
	return true
}

func (wl *WechatTokenLimit) WaitAllow(k string, duration time.Duration) (int64, error) {
	if wl.Allow(k) {
		return 0, nil
	}

	var cnt int64
	if duration/time.Hour > 24 {
		duration = time.Hour * 24
	}

	lctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	ticker := time.NewTicker(retryDuration)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			cnt++
			if wl.Allow(k) {
				return cnt * cast.ToInt64(retryDuration.Seconds()), nil
			}
		case <-lctx.Done():
			return cast.ToInt64(duration.Seconds()), lctx.Err()
		}
	}
}
