package svc

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhuud/go-library/svc/alarm"
)

const (
	DefaultReceiveType = "chat_id"
	DefaultReceiveId   = "oc_53b66a251e2a89ed74b4be3098262af5"
)

type Alarm struct{}

func (al *Alarm) SendLark(content string) {
	_ = alarm.Send(alarm.LarkMessage{
		ReceiveType: DefaultReceiveType,
		ReceiveId:   DefaultReceiveId,
		Content:     content,
	})
}

func (al *Alarm) SendLarkCtx(ctx context.Context, content string) {
	logx.WithContext(ctx).Error(content)
	al.SendLark(content)
}
