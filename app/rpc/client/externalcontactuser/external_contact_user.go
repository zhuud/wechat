// Code generated by goctl. DO NOT EDIT.
// Source: wechat.proto

package externalcontactuser

import (
	"context"

	"rpc/wechat"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Error                                   = wechat.Error
	ErrorResp                               = wechat.ErrorResp
	ExternalContactWayConclusion            = wechat.ExternalContactWayConclusion
	ExternalContactWayConclusionImage       = wechat.ExternalContactWayConclusionImage
	ExternalContactWayConclusionLink        = wechat.ExternalContactWayConclusionLink
	ExternalContactWayConclusionMiniprogram = wechat.ExternalContactWayConclusionMiniprogram
	ExternalContactWayConclusionText        = wechat.ExternalContactWayConclusionText
	ExternalContactWayData                  = wechat.ExternalContactWayData
	ExternalContactWayInfoResp              = wechat.ExternalContactWayInfoResp
	ExternalContactWayListReq               = wechat.ExternalContactWayListReq
	ExternalContactWayListResp              = wechat.ExternalContactWayListResp
	ExternalContactWayReq                   = wechat.ExternalContactWayReq
	ExternalGroupChatInfoReq                = wechat.ExternalGroupChatInfoReq
	ExternalGroupChatListReq                = wechat.ExternalGroupChatListReq
	ExternalUser                            = wechat.ExternalUser
	ExternalUserFollowUser                  = wechat.ExternalUserFollowUser
	ExternalUserFollowUserTag               = wechat.ExternalUserFollowUserTag
	ExternalUserFollowUserWechatChannel     = wechat.ExternalUserFollowUserWechatChannel
	ExternalUserIdReq                       = wechat.ExternalUserIdReq
	ExternalUserIdResp                      = wechat.ExternalUserIdResp
	ExternalUserInfo                        = wechat.ExternalUserInfo
	ExternalUserInfoOpt                     = wechat.ExternalUserInfoOpt
	ExternalUserInfoReq                     = wechat.ExternalUserInfoReq
	ExternalUserInfoResp                    = wechat.ExternalUserInfoResp
	ExternalUserProfile                     = wechat.ExternalUserProfile
	ExternalUserProfileItem                 = wechat.ExternalUserProfileItem
	ExternalUserProfileItemText             = wechat.ExternalUserProfileItemText
	SaveExternalContactWayResp              = wechat.SaveExternalContactWayResp
	UpdateExternalUserRemarkReq             = wechat.UpdateExternalUserRemarkReq
	UseridList                              = wechat.UseridList

	ExternalContactUser interface {
		GetExternalUserInfo(ctx context.Context, in *ExternalUserInfoReq, opts ...grpc.CallOption) (*ExternalUserInfoResp, error)
		GetExternalUserIdByUserId(ctx context.Context, in *ExternalUserIdReq, opts ...grpc.CallOption) (*ExternalUserIdResp, error)
		UpdateExternalUserRemark(ctx context.Context, in *UpdateExternalUserRemarkReq, opts ...grpc.CallOption) (*ErrorResp, error)
	}

	defaultExternalContactUser struct {
		cli zrpc.Client
	}
)

func NewExternalContactUser(cli zrpc.Client) ExternalContactUser {
	return &defaultExternalContactUser{
		cli: cli,
	}
}

func (m *defaultExternalContactUser) GetExternalUserInfo(ctx context.Context, in *ExternalUserInfoReq, opts ...grpc.CallOption) (*ExternalUserInfoResp, error) {
	client := wechat.NewExternalContactUserClient(m.cli.Conn())
	return client.GetExternalUserInfo(ctx, in, opts...)
}

func (m *defaultExternalContactUser) GetExternalUserIdByUserId(ctx context.Context, in *ExternalUserIdReq, opts ...grpc.CallOption) (*ExternalUserIdResp, error) {
	client := wechat.NewExternalContactUserClient(m.cli.Conn())
	return client.GetExternalUserIdByUserId(ctx, in, opts...)
}

func (m *defaultExternalContactUser) UpdateExternalUserRemark(ctx context.Context, in *UpdateExternalUserRemarkReq, opts ...grpc.CallOption) (*ErrorResp, error) {
	client := wechat.NewExternalContactUserClient(m.cli.Conn())
	return client.UpdateExternalUserRemark(ctx, in, opts...)
}
