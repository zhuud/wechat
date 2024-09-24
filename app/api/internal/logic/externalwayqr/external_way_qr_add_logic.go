package externalwayqr

import (
	"context"
	"encoding/json"
	"github.com/spf13/cast"
	"rpc/client/externalcontactway"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExternalWayQrAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 企微联系人二维码添加
func NewExternalWayQrAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExternalWayQrAddLogic {
	return &ExternalWayQrAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExternalWayQrAddLogic) ExternalWayQrAdd(req *types.ExternalContactWayRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	externalContactWayData := &externalcontactway.ExternalContactWayData{
		Type:          req.Type,
		Scene:         req.Scene,
		Style:         req.Style,
		Remark:        req.Remark,
		SkipVerify:    req.SkipVerify,
		State:         req.State,
		User:          req.User,
		Party:         req.Party,
		ExpiresIn:     cast.ToInt32(req.ExpiresIn),
		ChatExpiresIn: cast.ToInt32(req.ChatExpiresIn),
		Unionid:       req.UnionID,
	}

	conclusions := &externalcontactway.ExternalContactWayConclusion{}

	if req.ConclusionsText != "" {
		var textOfMessage externalcontactway.ExternalContactWayConclusionText
		json.Unmarshal([]byte(req.ConclusionsText), &textOfMessage)
		conclusions.Text = &textOfMessage
	}
	if req.ConclusionsLink != "" {
		var link externalcontactway.ExternalContactWayConclusionLink
		json.Unmarshal([]byte(req.ConclusionsLink), &link)
		conclusions.Link = &link
	}
	if req.ConclusionsImage != "" {
		var image externalcontactway.ExternalContactWayConclusionImage
		json.Unmarshal([]byte(req.ConclusionsImage), &image)
		conclusions.Image = &image
	}
	if req.ConclusionsMiniProgram != "" {
		var miniProgram externalcontactway.ExternalContactWayConclusionMiniprogram
		json.Unmarshal([]byte(req.ConclusionsMiniProgram), &miniProgram)
		conclusions.Miniprogram = &miniProgram
	}

	externalContactWayData.Conclusions = conclusions

	data, err := l.svcCtx.ExternalcontactwayRpc.CreateExternalContactWay(l.ctx, externalContactWayData)
	resp = &types.Response{
		Data: data,
	}

	return resp, err
}
