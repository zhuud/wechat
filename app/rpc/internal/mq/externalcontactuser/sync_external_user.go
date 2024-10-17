package externalcontactuser

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"rpc/internal/config"
	externalcontactuserlogic "rpc/internal/logic/externalcontactuser"
	"rpc/internal/svc"
	"rpc/internal/types"
	"rpc/model"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhuud/go-library/svc/kafka"
)

const syncExternalUserTopic = "5002,5003"

// 支持的所有操作类型
var syncStaffCallBackChangeTypeAll = []string{
	syncStaffCallBackChangeTypeCreateUser,
	syncStaffCallBackChangeTypeDelUser,
}

// 支持的所有操作类型
var syncUserCallBackChangeTypeAll = []string{
	syncUserCallBackChangeTypeAddC,
	syncUserCallBackChangeTypeAddHalfC,
	syncUserCallBackChangeTypeCDel,
	syncUserCallBackChangeTypeDelC,
}

// 操作类型 - 加c
var syncUserCallBackChangeTypeAdd = []string{
	syncUserCallBackChangeTypeAddC,
	syncUserCallBackChangeTypeAddHalfC,
}

// 操作类型 - 减c
var syncUserCallBackChangeTypeDel = []string{
	syncUserCallBackChangeTypeCDel,
	syncUserCallBackChangeTypeDelC,
}

const (
	syncUserCallBackEvent     = "change_external_contact" // 事件的类型，此时固定为change_external_contact
	syncUserChatCallBackEvent = "change_external_chat"    // 微信群
	syncUserTagCallBackEvent  = "change_external_tag"
	syncStaffCallBackEvent    = "change_contact"

	syncUserCallBackChangeTypeAddC        = `add_external_contact`      // 添加企业客户事件
	syncUserCallBackChangeTypeAddHalfC    = `add_half_external_contact` // 外部联系人免验证添加成员事件
	syncUserCallBackChangeTypeCDel        = `del_external_contact`      // 删除企业客户事件(员工主动)
	syncUserCallBackChangeTypeDelC        = `del_follow_user`           // 删除跟进成员事件(客户主动)
	syncUserCallBackChangeTypeMsgApproved = `msg_audit_approved`        //客户同意进行聊天内容存档

	syncStaffCallBackChangeTypeCreateUser = `create_user` // 添加企业客户事件
	syncStaffCallBackChangeTypeDelUser    = `delete_user`
)

func NewSyncExternalUserConsumer(kafkaConf kq.KqConf, svcCtx *svc.ServiceContext) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		kafkaConf.Topic = syncExternalUserTopic
		kafka.Consume(kafkaConf, cmd.Use, newSyncExternalUserConsumer(cmd.Context(), svcCtx))
	}
}

type syncExternalUserConsumer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func newSyncExternalUserConsumer(ctx context.Context, svcCtx *svc.ServiceContext) *syncExternalUserConsumer {
	return &syncExternalUserConsumer{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

type ExternalUserValue struct {
	types.TopicMessage
	Data Data `json:"data"`
}

type Data struct {
	CropName string             `json:"crop_name"`
	MsgData  types.ExternalData `json:"msg_data"`
}

func (s *syncExternalUserConsumer) Consume(ctx context.Context, key string, value string) error {
	//限流 并将消息写入延迟队列 todo
	if !s.svcCtx.WechatLimit.Allow(`external_user`) {
		return nil
	}
	time.Sleep(1 * time.Second)
	// do code ...
	spew.Dump(key, value)
	topicMessage := ExternalUserValue{}
	json.Unmarshal([]byte(value), &topicMessage)

	spew.Dump(topicMessage, 222)
	switch topicMessage.Topic {
	case config.MqTagQywxSyncUserCallback:
		s.handleUserCallBack(topicMessage)

	case config.MqTagQywxSyncStaffCallback:
	}

	return nil
}

// 客户消息回调
func (s *syncExternalUserConsumer) handleUserCallBack(message ExternalUserValue) {
	switch message.Data.MsgData.Event {
	//客户信息更新
	case syncUserCallBackEvent:
		s.handleExternalUser(message.Data)

		//群聊消息更新
	case syncUserChatCallBackEvent:

		//客户标签更新
	case syncUserTagCallBackEvent:

	}
}

// 处理用户信息
func (s *syncExternalUserConsumer) handleExternalUser(data Data) {
	msgData := data.MsgData
	if msgData.ExternalUserID == `` || msgData.UserID == `` {
		s.svcCtx.Alarm.SendLarkCtx(s.ctx, fmt.Sprintf("syncExternalUserCmd.exter.user.empty dui:%d error: %v 外部联系人，用户id为空", data, nil))
		return
	}

	//客户同意进行聊天内容存档
	if object.InArray(msgData.ChangeType, []string{syncUserCallBackChangeTypeMsgApproved}) {
		err := externalcontactuserlogic.NewSaveExternalUserLogic(s.ctx, s.svcCtx).UpdateChatAgreeStatus(msgData.ExternalUserID, msgData.UserID)
		if err != nil {
			s.svcCtx.Alarm.SendLarkCtx(s.ctx, fmt.Sprintf("syncExternalUserCmd.batchGetExternal.WaitAllow dui:%d error: %v 需要重新处理后续数据", data, err))
		}
		return
	}

	// 增加用户信息
	if object.InArray(msgData.ChangeType, []string{syncUserCallBackChangeTypeAddHalfC, syncUserCallBackChangeTypeAddC}) {
		//externalcontactuserlogic.NewSaveExternalUserLogic(s.ctx, s.svcCtx).SyncExternalUser(data)
	}

	// 编辑用户信息
	if object.InArray(msgData.ChangeType, []string{`edit_external_contact`, ``}) {

	}

	// 删除用户信息
	if object.InArray(msgData.ChangeType, []string{syncUserCallBackChangeTypeCDel, syncUserCallBackChangeTypeDelC}) {
		status := 0
		if msgData.ChangeType == syncUserCallBackChangeTypeCDel {
			status = model.CDelFollowUserStatus
		} else if msgData.ChangeType == syncUserCallBackChangeTypeDelC {
			status = model.FollowUserDelCStatus
		}
		err := externalcontactuserlogic.NewSaveExternalUserLogic(s.ctx, s.svcCtx).DeleteExternalUserFollow(data.CropName, data.MsgData, status)
		if err != nil {
			s.svcCtx.Alarm.SendLarkCtx(s.ctx, fmt.Sprintf("syncExternalUserCmd.DeleteExternalUserFollow.err dui:%d error: %v 需要重新处理后续数据", data, err))
		}
	}

	// 客户接替失败 todo
	if object.InArray(msgData.ChangeType, []string{`transfer_fail`}) {

	}

	//其他业务数据处理
	s.handleOtherData(msgData)
}

/**
 * 处理业务数据
 */
func (s *syncExternalUserConsumer) handleOtherData(msgData types.ExternalData) {
	if !object.InArray(msgData.ChangeType, []string{syncUserCallBackChangeTypeAddHalfC, syncUserCallBackChangeTypeAddC}) {
		return
	}

	//给活动发送私域添加消息
	s.sendDoActivityTaskMsg()
}

func (s *syncExternalUserConsumer) sendDoActivityTaskMsg() {

}

// 处理用户聊天信息
func (s *syncExternalUserConsumer) handleExternalUserChat(data Data) {

}

// 处理用户标签
func (s *syncExternalUserConsumer) handleExternalUserTag(data Data) {

}
