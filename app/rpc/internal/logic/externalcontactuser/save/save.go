package save

import (
	"context"
	"errors"
	"github.com/avast/retry-go"
	"rpc/internal/svc"
	"rpc/model"
	"time"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/response"
	"github.com/zeromicro/go-zero/core/logx"
)

type SaveExternalUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveExternalUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveExternalUserLogic {
	return &SaveExternalUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (s *SaveExternalUserLogic) Add(externalUser *response.ResponseExternalContact) error {
	if externalUser == nil || externalUser.ExternalContact == nil {
		return errors.New("externalUser is nil")
	}

	var (
		dbExternalUser *model.TbExternalUser
		err            error
	)
	_ = retry.Do(func() error {
		dbExternalUser, err = s.svcCtx.ModelExternalUser.FindOne(s.ctx, externalUser.ExternalContact.ExternalUserID)
		return err
	}, retry.Attempts(3))

	ts := time.Now().Local()
	if dbExternalUser != nil {
		err = s.svcCtx.ModelExternalUser.Update(s.ctx, &model.TbExternalUser{
			ExternalUserid: externalUser.ExternalContact.ExternalUserID,
			Unionid:        externalUser.ExternalContact.UnionID,
			Type:           uint64(externalUser.ExternalContact.Type),
			Name:           externalUser.ExternalContact.Name,
			Avatar:         externalUser.ExternalContact.Avatar,
			Gender:         uint64(externalUser.ExternalContact.Gender),
			CorpName:       externalUser.ExternalContact.CorpName,
			CorpFullName:   externalUser.ExternalContact.CorpFullName,
			Position:       externalUser.ExternalContact.Position,
			Status:         model.TbExternalUserNormalStatus,
			CreatedAt:      ts,
			UpdatedAt:      ts,
		})
	} else {
		_, err = s.svcCtx.ModelExternalUser.Insert(s.ctx, &model.TbExternalUser{
			ExternalUserid: externalUser.ExternalContact.ExternalUserID,
			Unionid:        externalUser.ExternalContact.UnionID,
			Type:           uint64(externalUser.ExternalContact.Type),
			Name:           externalUser.ExternalContact.Name,
			Avatar:         externalUser.ExternalContact.Avatar,
			Gender:         uint64(externalUser.ExternalContact.Gender),
			CorpName:       externalUser.ExternalContact.CorpName,
			CorpFullName:   externalUser.ExternalContact.CorpFullName,
			Position:       externalUser.ExternalContact.Position,
			Status:         model.TbExternalUserNormalStatus,
			CreatedAt:      ts,
			UpdatedAt:      ts,
		})
	}

	if err != nil {
		return err
	}

	return nil
}
