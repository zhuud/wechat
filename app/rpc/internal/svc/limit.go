package svc

import (
	"fmt"
	"sync"

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
		redis    *redis.Redis
	}
	WechatLimiter [2]*limit.TokenLimiter
)

var (
	wl       *WechatTokenLimit
	wlOnce   sync.Once
	keyQuota = map[string][2]int{
		"all":          {minuteQuota, hourQuota},
		"external_tag": {5000, 50000},
	}
)

func NewTokenWechatLimit(r *redis.Redis) *WechatTokenLimit {

	wlOnce.Do(func() {
		wl = &WechatTokenLimit{
			limiters: make(map[string]*WechatLimiter),
			redis:    r,
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
