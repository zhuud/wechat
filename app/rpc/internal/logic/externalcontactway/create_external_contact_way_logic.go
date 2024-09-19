package externalcontactwaylogic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/contactWay/request"
	msgTplgReq "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/messageTemplate/request"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
	"rpc/internal/svc"
	"rpc/model"
	"rpc/wechat"
	"strings"
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

	if in.Type == 0 {
		return nil, errors.New("参数错误-type")
	}
	if in.Scene == 0 {
		return nil, errors.New("参数错误-scene")
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

	conclusionType := ""
	conclusionContent := ""

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
			conclusionType = "text"
			jsonData, _ := json.Marshal(options.Conclusions.Text)
			conclusionContent = cast.ToString(jsonData)
		}
		if in.Conclusions.Link != nil {
			options.Conclusions.Link = &msgTplgReq.Link{
				Title:  in.Conclusions.Link.Title,
				PicURL: in.Conclusions.Link.Picurl,
				Desc:   in.Conclusions.Link.Desc,
				URL:    in.Conclusions.Link.Url,
			}
			conclusionType = "link"
			jsonData, _ := json.Marshal(options.Conclusions.Link)
			conclusionContent = cast.ToString(jsonData)
		}
		if in.Conclusions.Image != nil {
			options.Conclusions.Image = &msgTplgReq.Image{
				MediaID: in.Conclusions.Image.MediaId,
				PicURL:  "",
			}
			conclusionType = "image"
			jsonData, _ := json.Marshal(options.Conclusions.Image)
			conclusionContent = cast.ToString(jsonData)
		}
		if in.Conclusions.Miniprogram != nil {
			options.Conclusions.MiniProgram = &msgTplgReq.MiniProgram{
				Title:      in.Conclusions.Miniprogram.Title,
				PicMediaID: in.Conclusions.Miniprogram.PicMediaId,
				AppID:      in.Conclusions.Miniprogram.Appid,
				Page:       in.Conclusions.Miniprogram.Page,
			}
			conclusionType = "miniprogram"
			jsonData, _ := json.Marshal(options.Conclusions.MiniProgram)
			conclusionContent = cast.ToString(jsonData)
		}
	}

	// 调用企微配置客户联系「联系我」方式接口
	resAdd, err := l.svcCtx.WeCom.WithCorp("yx").ContactWay.Add(l.ctx, options)
	if err != nil {
		l.Logger.Error("ContactWay_Add_Err", err)
		return nil, err
	}

	if resAdd != nil && resAdd.ErrCode != 0 {
		l.Logger.Error("ContactWay_Add_Result_Err", resAdd)
		return nil, errors.New(resAdd.ErrMsg)
	}

	// 本地结构化
	userServiceQrcode := &model.UserServiceQrcode{
		ConfigId:      resAdd.ConfigID,
		Type:          cast.ToInt64(in.Type),
		Scene:         cast.ToInt64(in.Scene),
		Style:         cast.ToInt64(in.Style),
		Remark:        in.Remark,
		SkipVerify:    cast.ToInt64(in.SkipVerify),
		State:         in.State,
		QrCode:        resAdd.QRCode,
		User:          strings.Join(in.User, ","),
		Party:         strings.Join(cast.ToStringSlice(interface{}(in.Party)), ","),
		IsTemp:        cast.ToInt64(in.IsTemp),
		ExpiresIn:     cast.ToInt64(in.ExpiresIn),
		ChatExpiresIn: cast.ToInt64(in.ChatExpiresIn),
		Unionid:       in.Unionid,
		IsExclusive:   cast.ToInt64(in.IsExclusive),
		Status:        1,
	}
	insertRes, insertErr := l.svcCtx.ModelUserServiceQrcodeModel.Insert(l.ctx, userServiceQrcode)
	if insertErr != nil {
		l.Logger.Error("ModelUserServiceQrcodeModel_Insert_Err", insertErr)
		return nil, insertErr
	}

	if insertRes != nil {
		lastInsertId, _ := insertRes.LastInsertId()
		if lastInsertId > 0 {
			userServiceQrcodeConclusion := &model.TbUserServiceQrcodeConclusions{
				UserServiceQcCodeId: lastInsertId,
				Type:                conclusionType,
				Content:             conclusionContent,
				Status:              1,
			}

			_, conclusionInsertErr := l.svcCtx.ModelUserServiceQrcodeConclusion.Insert(l.ctx, userServiceQrcodeConclusion)
			if conclusionInsertErr != nil {
				l.Logger.Error("ModelUserServiceQrcodeConclusion_Insert_Err", conclusionInsertErr)
				return nil, conclusionInsertErr
			}
		}
	}

	result := &wechat.SaveExternalContactWayResp{
		ConfigId: resAdd.ConfigID,
		QrCode:   resAdd.QRCode,
	}

	return result, nil
}
