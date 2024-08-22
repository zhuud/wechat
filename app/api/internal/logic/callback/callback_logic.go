package callback

import (
    "api/internal/svc"
    "api/internal/types"
    "context"

    "github.com/zeromicro/go-zero/core/logx"
)

type CallbackLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

// 企微回调处理
func NewCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CallbackLogic {
    return &CallbackLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *CallbackLogic) Callback(req *types.CallbackReq) (resp *types.Response, err error) {
    // todo: add your logic here and delete this line

    return &types.Response{
        Msg: "success",
    }, nil
}
