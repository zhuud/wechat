package svc

import (
	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

const (
	minutePeriod = 60
	minuteQuota  = 5000

	hourPeriod = 3600
	hourQuota  = 150000

	minuteKeyPreFix = "wechat:period:second:limit:"
	hourKeyPreFix   = "wechat:period:hour:limit:"
)

type WechatPeriodLimit struct {
	MinuteLimit *limit.PeriodLimit
	HourLimit   *limit.PeriodLimit
}

func NewWechatPeriodLimit(redis *redis.Redis) *WechatPeriodLimit {
	ml := limit.NewPeriodLimit(minutePeriod, minuteQuota, redis, minuteKeyPreFix)
	hl := limit.NewPeriodLimit(hourPeriod, hourQuota, redis, hourKeyPreFix)

	return &WechatPeriodLimit{
		MinuteLimit: ml,
		HourLimit:   hl,
	}
}

func (wl *WechatPeriodLimit) Take(key string) bool {
	secondCode, secondErr := wl.MinuteLimit.Take(key)
	hourCode, hourErr := wl.HourLimit.Take(key)
	if secondErr != nil {
		logx.Errorf(`WechatLimit secondErr: %v`, secondErr)
	}
	if hourErr != nil {
		logx.Errorf(`WechatLimit hourErr: %v`, secondErr)
	}
	if secondCode == limit.OverQuota {
		logx.Infof(`WechatLimit secondKey: %s`, key)
		return false
	}
	if hourCode == limit.OverQuota {
		logx.Infof(`WechatLimit hourKey: %s`, key)
		return false
	}
	return true
}
