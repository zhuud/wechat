package callback

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cast"
	"rpc/internal/config"
	externalcontactwaylogic "rpc/internal/logic/externalcontactway"

	//externalcontactwaylogic "rpc/internal/logic/externalcontactway"
	"rpc/internal/svc"
	"rpc/model"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhuud/go-library/svc/conf"
	"github.com/zhuud/go-library/svc/kafka"
	"github.com/zhuud/go-library/utils"
)

type CallbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CallbackLogic {
	return &CallbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

var syncUserCallBackEvent = config.ChangeExternalContactEvent // 事件的类型，此时固定为change_external_contact
var syncStaffCallBackEvent = config.StaffChangeContactEvent

var syncUserCallBackChangeTypeAddC = config.AddExternalContactChangeType         // 添加企业客户事件
var syncUserCallBackChangeTypeAddHalfC = config.AddHalfExternalContactChangeType // 外部联系人免验证添加成员事件
var syncUserCallBackChangeTypeCDel = config.DelExternalContactChangeType         // 删除企业客户事件(员工主动)
var syncUserCallBackChangeTypeDelC = config.DelFollowUserChangeType              // 删除跟进成员事件(客户主动)

var syncStaffCallBackChangeTypeCreateUser = config.StaffCreateUser // 添加企业客户事件
var syncStaffCallBackChangeTypeUpdateUser = config.StaffUpdateUser // 编辑企业客户事件
var syncStaffCallBackChangeTypeDelUser = config.StaffDeleteUser

// 支持的所有操作类型
var syncStaffCallBackChangeTypeAll = []string{
	syncStaffCallBackChangeTypeCreateUser,
	syncStaffCallBackChangeTypeUpdateUser,
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

var wbToPlatformTest = map[string]string{
	"wb:1:105632160": config.ContactWayBizTypeWeChatAddUser, // 微办渠道码上的state对应的平台业务
}

var wbToPlatform = map[string]string{
	"wb:1:105710252": config.ContactWayBizTypeWeChatAddUser, // 微办渠道码上的state对应的平台业务
}

// 外部联系人
func (l *CallbackLogic) HandleSyncUserCallback(data map[string]interface{}) bool {
	l.Info(l.ctx, "handleSyncUserCallback", data)

	cropName := data[`crop_name`]
	msgData := data[`msg_data`].(map[string]any)

	event := cast.ToString(msgData[`Event`])
	changeType := cast.ToString(msgData[`ChangeType`])
	externalUserID := cast.ToString(msgData[`ExternalUserID`])
	staffId := cast.ToString(msgData[`UserID`])
	//welcomeCode := cast.ToString(msgData[`WelcomeCode`])

	if event != syncUserCallBackEvent {
		l.svcCtx.Alarm.SendLarkCtx(l.ctx, fmt.Sprintf("handleSyncUserCallback_Event error: %v ", data))
		l.Info(l.ctx, `handleSyncUserCallback_Event error`, event, syncUserCallBackEvent, data)
		return false
	}

	if !utils.ArrayIn(changeType, syncUserCallBackChangeTypeAll) {
		l.svcCtx.Alarm.SendLarkCtx(l.ctx, fmt.Sprintf("handleSyncUserCallback_changeType不予处理 error: %v ", data))
		l.Info(l.ctx, `handleSyncUserCallback_changeType`, data)
		return false
	}

	if externalUserID == `` || staffId == `` {
		l.svcCtx.Alarm.SendLarkCtx(l.ctx, fmt.Sprintf("handleSyncUserCallback_externalUserID_staffId_state_nil error: %v ", data))
		l.Info(l.ctx, `handleSyncUserCallback_externalUserID_staffId_state_nil`, data)
		return false
	}

	if cropName != `youxiang` {
		l.svcCtx.Alarm.SendLarkCtx(l.ctx, fmt.Sprintf("handleSyncUserCallback_cropName_not_youxiang error: %v ", data))
		l.Info(l.ctx, `handleSyncUserCallback_cropName_not_youxiang`, data)
		return false
	}

	// 用户详情
	userInfo, _ := l.svcCtx.ModelExternalUser.FindOne(l.ctx, externalUserID)
	unionId := userInfo.Unionid
	if unionId == `` {
		l.svcCtx.Alarm.SendLarkCtx(l.ctx, fmt.Sprintf("handleSyncUserCallback_union_id_nil error: %v ", data))
		l.Info(l.ctx, `handleSyncUserCallback_union_id_nil`, data, userInfo)
		return false
	}

	// 加c
	if utils.ArrayIn(changeType, syncUserCallBackChangeTypeAdd) {
		state := cast.ToString(msgData[`State`])
		if state == `` {
			l.svcCtx.Alarm.SendLarkCtx(l.ctx, fmt.Sprintf("handleSyncUserCallback_state_empty error: %v ", data))
			l.Info(l.ctx, `handleSyncUserCallback_state_empty`, data, userInfo)
			return false
		}

		// 加c二维码  码上透传参数state设置 （因为state只支持30字符透传  ）
		// 1. 业务类型 定义目录 wechat/app/rpc/internal/config/const.go
		// 2. 透传参数定义 json
		// 3. state格式 eg "state": "99_9999#1#01234567890123456789"  或者 "99_9999#0#{"hid":15731}"
		// 	# 分割符号
		//	前7位为渠道来源
		//	9位是否使用临时缓存
		//	后20位为业务信息, 如果信息20位存不下，使用临时缓存将信息生成唯一key add_contact_way_%s_%s，回调时候再从缓存获取
		//
		// 动态加c码
		//	wecom.ExternalContactBusiness.AddContactWay（正常直接传透传json就好，业务里长度大于30自动转换走缓存）
		// 静态码
		//	微办后台生成  （注：如果静态码也有回调处理的相关逻辑 state设置格式要同动态码 "99_9999#0#{"hid":15731}" 定义好让运营设置 ）
		//
		// 回调逻辑处理
		// 新策略相关信息透传
		if bizType, callback, flag := l.handleState(state, msgData); flag {
			// 分发到各业务里
			l.distributeBiz(bizType, callback, userInfo, msgData)

		} else {
			// 微办平台生成的码的兼容转发 callback信息里只有state
			wbToPlatformData := wbToPlatform
			if !conf.IsProd() {
				wbToPlatformData = wbToPlatformTest
			}
			if bizType, ok := wbToPlatformData[state]; ok {
				l.distributeBiz(bizType, "", userInfo, msgData)
			} else {
				l.distributeBiz(config.ContactWayBizTypeChainOldBringNewC, state, userInfo, msgData)
			}

		}
	}

	// 减c
	if utils.ArrayIn(changeType, syncUserCallBackChangeTypeDel) {
		// todo: add your logic here and delete this line
	}

	return true
}

func (l *CallbackLogic) distributeBiz(bizType, callback string, userInfo *model.TbExternalUser, data map[string]any) {

	l.Info(l.ctx, "distributeBiz.distribute", data)

	changeType := cast.ToString(data[`ChangeType`])
	staffId := cast.ToString(data[`UserID`])
	// welcomeCode := cast.ToString(data[`WelcomeCode`])

	// TODO  后续支持加欢迎语
	// TODO 后续可以统计员工已加c数量 快到达5w上限 报警
	topicData := map[string]any{
		`biz_type`:    bizType,    // 申请的业务type
		`callback`:    callback,   // 透传的参数
		`user_info`:   userInfo,   // c用户信息
		`staff_id`:    staffId,    // 员工ID
		`change_type`: changeType, // 员工加c 还是c加员工
	}
	tj, _ := jsoniter.MarshalToString(topicData)

	spew.Dump("转发消息：", tj)
	_ = kafka.Push(l.ctx, config.MqTagJoinCCallback, topicData)
}

func (l *CallbackLogic) handleState(state string, data map[string]interface{}) (bizType, callback string, flag bool) {
	stateList := make([]string, 0)
	// "state": "99_9999#1#01234567890123456789"  或者 "99_9999#0#{"hid":15731}"
	// 充分校验下 防止正常渠道里带#号
	if !strings.Contains(state, "#") {
		return state, "", false
	}
	if stateList = strings.Split(state, "#"); len(stateList) != 3 {
		return state, "", false
	}
	if !utils.ArrayIn(stateList[1], []string{"0", "1"}) {
		return state, "", false
	}
	bizType = stateList[0]
	callback = stateList[2]
	// 使用缓存的从缓存读取
	if userCache := stateList[1]; userCache == "1" {
		externalContactWayLogic := externalcontactwaylogic.NewCreateExternalContactWayLogic(l.ctx, l.svcCtx)
		callbackNew, err := externalContactWayLogic.GetContactWayCallback(bizType, callback)
		if err != nil {
			callbackNew, err = externalContactWayLogic.GetContactWayCallback(bizType, callback)
		}
		if err != nil {
			l.svcCtx.Alarm.SendLarkCtx(l.ctx, fmt.Sprintf("handleState_GetContactWayCallback error: %s ", state))
			l.Info(l.ctx, `handleState_GetContactWayCallback_error`, data)
		}
		callback = callbackNew
	}

	return bizType, callback, true
}
