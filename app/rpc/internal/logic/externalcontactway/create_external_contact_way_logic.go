package externalcontactwaylogic

import (
	"context"
	"errors"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/contactWay/request"
	"github.com/spf13/cast"
	"rpc/internal/svc"
	"rpc/wechat"

	msgTplgReq "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/messageTemplate/request"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateExternalContactWayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateExternalContactWayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateExternalContactWayLogic {
	return &CreateExternalContactWayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateExternalContactWayLogic) CreateExternalContactWay(in *wechat.ExternalContactWayData) (*wechat.SaveExternalContactWayResp, error) {
	// todo: add your logic here and delete this line
	res := &wechat.SaveExternalContactWayResp{}

	if in.Type == 0 {
		return res, errors.New("参数错误-type")
	}
	if in.Scene == 0 {
		return res, errors.New("参数错误-scene")
	}

	options := &request.RequestAddContactWay{
		Type:          cast.ToInt(in.Type),
		Scene:         cast.ToInt(in.Scene),
		Style:         cast.ToInt(in.Style),
		Remark:        in.Remark,
		SkipVerify:    in.SkipVerify,
		State:         in.State,
		User:          in.User,
		Party:         cast.ToIntSlice(interface{}(in.Party)),
		IsTemp:        in.IsTemp,
		ExpiresIn:     cast.ToInt(in.ExpiresIn),
		ChatExpiresIn: cast.ToInt(in.ChatExpiresIn),
		UnionID:       in.Unionid,
	}

	if in.Conclusions != nil {
		options.Conclusions = &request.Conclusions{
			Text: &msgTplgReq.TextOfMessage{Content: ""},
			Image: &msgTplgReq.Image{
				MediaID: "",
				PicURL:  "",
			},
			Link: &msgTplgReq.Link{
				Title:  "",
				PicURL: "",
				Desc:   "",
				URL:    "",
			},
			MiniProgram: &msgTplgReq.MiniProgram{
				Title:      "",
				PicMediaID: "",
				AppID:      "",
				Page:       "",
			},
		}

		if in.Conclusions.Text != nil {
			options.Conclusions.Text = &msgTplgReq.TextOfMessage{Content: in.Conclusions.Text.Content}
		}
		if in.Conclusions.Link != nil {
			options.Conclusions.Link = &msgTplgReq.Link{
				Title:  in.Conclusions.Link.Title,
				PicURL: in.Conclusions.Link.Picurl,
				Desc:   in.Conclusions.Link.Desc,
				URL:    in.Conclusions.Link.Url,
			}
		}
		if in.Conclusions.Image != nil {
			options.Conclusions.Image = &msgTplgReq.Image{
				MediaID: in.Conclusions.Image.MediaId,
				PicURL:  "",
			}
		}
		if in.Conclusions.Miniprogram != nil {
			options.Conclusions.MiniProgram = &msgTplgReq.MiniProgram{
				Title:      in.Conclusions.Miniprogram.Title,
				PicMediaID: in.Conclusions.Miniprogram.PicMediaId,
				AppID:      in.Conclusions.Miniprogram.Appid,
				Page:       in.Conclusions.Miniprogram.Page,
			}
		}
	}

	l.svcCtx.WeCom.WithCorp("yx").ContactWay.Add(l.ctx, options)

	return res, nil
}
