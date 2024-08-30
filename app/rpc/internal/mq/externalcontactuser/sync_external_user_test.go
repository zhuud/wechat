package externalcontactuser

import (
	"context"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/zeromicro/go-zero/core/proc"
	"github.com/zhuud/go-library/svc/kafka"
)

func Test_Mq(t *testing.T) {
	err := kafka.Push(context.Background(), "5002", map[string]interface{}{
		"uid": 11,
	})
	proc.Shutdown()
	spew.Dump(err)
}
