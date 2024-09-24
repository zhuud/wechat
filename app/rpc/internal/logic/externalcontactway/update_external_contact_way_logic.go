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

type UpdateExternalContactWayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateExternalContactWayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateExternalContactWayLogic {
	return &UpdateExternalContactWayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateExternalContactWayLogic) UpdateExternalContactWay(in *wechat.ExternalContactWayData) (*wechat.SaveExternalContactWayResp, error) {
	// todo: add your logic here and delete this line

	if in.ConfigId == "" {
		return nil, errors.New("参数错误-config_id")
	}

	options := &request.RequestUpdateContactWay{
		ConfigID:      in.ConfigId,
		Style:         cast.ToInt(in.Style),
		Remark:        in.Remark,
		SkipVerify:    in.SkipVerify,
		State:         in.State,
		User:          in.User,
		Party:         cast.ToIntSlice(interface{}(in.Party)),
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

	// 调用企微更新企业已配置的「联系我」方式接口
	resUpdate, err := l.svcCtx.WeCom.WithCorp("yx").ContactWay.Update(l.ctx, options)
	if err != nil {
		l.Logger.Error("ContactWay_Update_Err", err)
		return nil, err
	}

	if resUpdate != nil && resUpdate.ErrCode != 0 {
		l.Logger.Error("ContactWay_Update_Result_Err", resUpdate)
		return nil, errors.New(resUpdate.ErrMsg)
	}

	// 本地结构化-更新
	qrcodeInfo, findErr := l.svcCtx.ModelUserServiceQrcodeModel.FindOneByConfigId(l.ctx, in.ConfigId)
	if findErr != nil {
		l.Logger.Error("ModelUserServiceQrcodeModel_FindOneByConfigId_Err", findErr)
		return nil, findErr
	}

	if qrcodeInfo != nil {
		updateUserServiceQrcode := &model.UserServiceQrcode{
			ConfigId:      in.ConfigId,
			Type:          qrcodeInfo.Type,
			Scene:         qrcodeInfo.Scene,
			Style:         cast.ToInt64(in.Style),
			Remark:        in.Remark,
			SkipVerify:    cast.ToInt64(in.SkipVerify),
			State:         in.State,
			QrCode:        qrcodeInfo.QrCode,
			User:          strings.Join(in.User, ","),
			Party:         strings.Join(cast.ToStringSlice(interface{}(in.Party)), ","),
			IsTemp:        qrcodeInfo.IsTemp,
			ExpiresIn:     cast.ToInt64(in.ExpiresIn),
			ChatExpiresIn: cast.ToInt64(in.ChatExpiresIn),
			Unionid:       in.Unionid,
			IsExclusive:   qrcodeInfo.IsExclusive,
			Status:        qrcodeInfo.Status,
		}
		updateErr := l.svcCtx.ModelUserServiceQrcodeModel.Update(l.ctx, updateUserServiceQrcode)
		if updateErr != nil {
			l.Logger.Error("ModelUserServiceQrcodeModel_Update_Err", updateErr)
			return nil, updateErr
		}

		//结束语
		l.svcCtx.ModelUserServiceQrcodeConclusion.Delete(l.ctx, qrcodeInfo.Id)

		userServiceQrcodeConclusion := &model.TbUserServiceQrcodeConclusions{
			UserServiceQcCodeId: cast.ToInt64(qrcodeInfo.Id),
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

	result := &wechat.SaveExternalContactWayResp{
		ConfigId: qrcodeInfo.ConfigId,
		QrCode:   qrcodeInfo.QrCode,
	}

	return result, nil
}
