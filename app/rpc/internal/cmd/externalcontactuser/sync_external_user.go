package externalcontactuser

import (
	"context"
	"errors"
	"fmt"
	"time"

	"rpc/internal/config"
	"rpc/internal/logic/externalcontactuser/save"
	"rpc/internal/svc"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/response"
	"github.com/avast/retry-go"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/fx"
	"github.com/zeromicro/go-zero/core/logx"
)

// go run main.go -f etc/config.local.yaml CmdSyncExternalUser yx

func NewSyncExternalUserCmd(svcCtx *svc.ServiceContext) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return newSyncExternalUserCmd(cmd.Context(), svcCtx).Do(args)
	}
}

type syncExternalUserCmd struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func newSyncExternalUserCmd(ctx context.Context, svcCtx *svc.ServiceContext) *syncExternalUserCmd {
	return &syncExternalUserCmd{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (s *syncExternalUserCmd) Do(args []string) error {
	crop := config.CropYx
	if len(args) > 0 {
		crop = args[0]
	}
	saveLogic := save.NewSaveExternalUserLogic(s.ctx, s.svcCtx)

	// 获取配置了客户联系功能的成员列表
	generateFunc := func(source chan<- any) {
		userIdList, err := s.getFollowUserList()
		// TODO 上线打印
		// s.Infof("syncExternalUserCmd.getFollowUserList.userIdList: %v", userIdList)
		if err != nil {
			s.svcCtx.Alarm.SendLarkCtx(s.ctx, fmt.Sprintf("syncExternalUserCmd.GetFollowUsers error: %v", err))
		}

		// TODO 测试 只跑一个
		source <- "kanyaYang"
		return

		for _, userId := range userIdList {
			source <- userId
		}
	}

	// 获取外部用户信息
	getExternalUserFunc := func(item any) any {
		userId, _ := item.(string)

		externalUserList, err := s.batchGetExternal(crop, []string{userId})
		s.Infof("syncExternalUserCmd.batchGetExternal.userId: %s, userList: %v", userId, externalUserList)
		if err != nil {
			s.svcCtx.Alarm.SendLarkCtx(s.ctx, fmt.Sprintf("syncExternalUserCmd.BatchGet error list: \n %v", err))
			return externalUserList
		}

		return externalUserList
	}

	// 保存外部用户信息
	saveExternalUserFunc := func(item any) {
		externalUserList, _ := item.([]*response.ResponseExternalContact)

		for _, externalUser := range externalUserList {
			s.Infof("syncExternalUserCmd.saveExternalUser info: %s", externalUser.ExternalContact.ExternalUserID)
			if err := saveLogic.Save(crop, externalUser); err != nil {
				s.svcCtx.Alarm.SendLarkCtx(s.ctx, fmt.Sprintf("syncExternalUserCmd.saveExternalUser error: %v", err))
			}
		}
	}

	// 主程序执行
	fx.From(generateFunc).
		Map(getExternalUserFunc).
		Parallel(saveExternalUserFunc, fx.WithWorkers(64))

	return nil
}

func (s *syncExternalUserCmd) batchGetExternal(crop string, userIdList []string) ([]*response.ResponseExternalContact, error) {
	var (
		err              []error
		externalUserList []*response.ResponseExternalContact
		cursor           = ""
		// TODO 测试
		limit = 1
		// limit            = 100
		maxSize = 100000
		size    = 0
	)

	for {
		// 一个user最多10000用户 正常不会循环这么多次
		size++
		if size > maxSize {
			s.svcCtx.Alarm.SendLarkCtx(s.ctx, fmt.Sprintf("syncExternalUserCmd.batchGetExternal.maxSize userIdList: %v, cursor:%s, error: 分页获取循坏异常 有可能死循环", userIdList, cursor))
			break
		}

		// 频控限制
		dur, lerr := s.svcCtx.WechatLimit.WaitAllow("external_user", time.Hour*2)
		if lerr != nil {
			s.svcCtx.Alarm.SendLarkCtx(s.ctx, fmt.Sprintf("syncExternalUserCmd.batchGetExternal.WaitAllow userIdList: %v, cursor:%s, dur:%d, error: %v 微信频控限制需要重新处理后续数据", userIdList, cursor, dur, lerr))
			break
		}

		// 获取
		doerr := retry.Do(func() error {
			externalUserData, err := s.svcCtx.WeCom.WithCorp(crop).ExternalUser.BatchGet(s.ctx, userIdList, cursor, limit)
			if err != nil {
				return err
			}
			if externalUserData.ErrCode != 0 {
				return errors.New(fmt.Sprintf("userIdList: %v, cursor:%s, error: %v", userIdList, cursor, externalUserData.ErrMsg))
			}

			externalUserList = append(externalUserList, externalUserData.ExternalContactList...)
			cursor = externalUserData.NextCursor

			return nil
		}, retry.Attempts(1))

		err = append(err, doerr)

		// TODO 测试一条
		cursor = ""

		if len(cursor) == 0 {
			break
		}
	}

	return externalUserList, errors.Join(err...)
}

func (s *syncExternalUserCmd) getFollowUserList() ([]string, error) {
	userIdList := make([]string, 0)

	err := retry.Do(func() error {
		userList, err := s.svcCtx.WeCom.WithCorp(config.CropYx).ExternalUser.GetFollowUsers(s.ctx)
		if err != nil {
			return err
		}
		if userList.ErrCode != 0 {
			return errors.New(userList.ErrMsg)
		}
		for _, item := range userList.FollowUser {
			userIdList = append(userIdList, item)
		}
		return nil
	})

	return userIdList, err
}
