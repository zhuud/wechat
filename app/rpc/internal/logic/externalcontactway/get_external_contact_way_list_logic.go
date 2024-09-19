package externalcontactwaylogic

import (
	"context"
	"fmt"

	"rpc/internal/svc"
	"rpc/wechat"

	contactWayRequest "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/contactWay/request"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetExternalContactWayListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetExternalContactWayListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExternalContactWayListLogic {
	return &GetExternalContactWayListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetExternalContactWayListLogic) GetExternalContactWayList(in *wechat.ExternalContactWayListReq) (*wechat.ExternalContactWayListResp, error) {
	// todo: add your logic here and delete this line

	fmt.Println("in", in)
	params := &contactWayRequest.RequestListContactWay{
		Cursor:    in.Cursor,
		Limit:     int(in.Limit),
		StartTime: int64(in.StartTime),
		EndTime:   int64(in.EndTime),
	}

	list, err := l.svcCtx.WeCom.WithCorp("yx").ContactWay.List(context.Background(), params)

	if err != nil {
		return nil, err
	}

	fmt.Println(list)

	externalContactWayReqList := make([]*wechat.ExternalContactWayReq, 0)
	if len(list.ContactWayIDs) > 0 {
		for _, item := range list.ContactWayIDs {
			externalContactWayReqList = append(externalContactWayReqList, &wechat.ExternalContactWayReq{
				ConfigId: item.ConfigID,
			})
		}
	}

	fmt.Println("externalContactWayReqList", externalContactWayReqList)

	return &wechat.ExternalContactWayListResp{
		ContactWay: externalContactWayReqList,
	}, nil
}
