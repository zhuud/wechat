package externalcontactuser

import (
	"context"
	"encoding/json"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"rpc/internal/config"
	externalcontactuserlogic "rpc/internal/logic/externalcontactuser"
	"rpc/internal/svc"
	"rpc/internal/types"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhuud/go-library/svc/kafka"
)

func NewSyncExternalUserConsumer(kafkaConf kq.KqConf, svcCtx *svc.ServiceContext) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		kafkaConf.Topic = "5002,5003"
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
	Data types.ExternalData `json:"data"`
}

type Data struct {
	ToUserName     string `json:"ToUserName"`
	FromUserName   string `json:"FromUserName"`
	CreateTime     string `json:"CreateTime"`
	MsgType        string `json:"MsgType"`
	Event          string `json:"Event"`
	ChangeType     string `json:"ChangeType"`
	UserID         string `json:"UserID"`
	ExternalUserID string `json:"ExternalUserID"`
}

func (s *syncExternalUserConsumer) Consume(ctx context.Context, key string, value string) error {

	// do code ...
	spew.Dump(key, value)
	topicMessage := ExternalUserValue{}
	json.Unmarshal([]byte(value), &topicMessage)

	spew.Dump(topicMessage, 222)
	switch topicMessage.Topic {
	case config.MqTagQywxSyncUserCallback:
		s.handleUserCallBack(topicMessage)
	}

	return nil
}

// 客户消息回调
func (s *syncExternalUserConsumer) handleUserCallBack(message ExternalUserValue) {
	switch message.Data.Event {
	//客户信息更新
	case `change_external_contact`:
		s.handleExternalUser(message.Data)

		//群聊消息更新
	case `change_external_chat`:

		//客户标签更新
	case `change_external_tag`:

	}
}

// 处理用户信息
func (s *syncExternalUserConsumer) handleExternalUser(data types.ExternalData) {
	// 增加用户信息
	if object.InArray(data.ChangeType, []string{`add_half_external_contact`, `add_external_contact`}) {
		externalcontactuserlogic.NewSaveExternalUserLogic(s.ctx, s.svcCtx).SyncExternalUser(data)
	}

	// 编辑用户信息
	if object.InArray(data.ChangeType, []string{`edit_external_contact`, ``}) {

	}

	// 删除用户信息
	if object.InArray(data.ChangeType, []string{`del_external_contact`, `del_follow_user`}) {
		externalcontactuserlogic.NewSaveExternalUserLogic(s.ctx, s.svcCtx).DeleteExternalUserFollow(data.ExternalUserID, data.UserID, data.ChangeType)
	}

	// 客户接替失败 todo
	if object.InArray(data.ChangeType, []string{`transfer_fail`}) {

	}
}

// 处理用户聊天信息
func (s *syncExternalUserConsumer) handleExternalUserChat(data Data) {

}

// 处理用户标签
func (s *syncExternalUserConsumer) handleExternalUserTag(data Data) {

}
