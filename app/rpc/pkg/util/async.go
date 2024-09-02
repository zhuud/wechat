package util

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"sync"
)

type Async struct {
	wg sync.WaitGroup
}

func (t *Async) Go(ctx context.Context, f func(arg []any), arg ...any) {
	// arg 是循环里的参数 如果不传会导致被for 覆盖

	logName := ``
	if len(arg) > 0 {
		logName, _ = arg[0].(string)
	}

	t.wg.Add(1)

	go func() {
		defer func() {
			if err := recover(); err != nil {
				logc.Info(ctx, `async-并行获取信息出错`, err, logName)
			}
			t.wg.Done()
		}()
		//goCost := logx.WithDuration(100)
		f(arg)
		//goCost.Info(`func-async -` + logName)
	}()
}

func (t *Async) Wait() {
	t.wg.Wait()
}
