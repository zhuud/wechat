package config

const (
	SoyoungCorp  = "soyoung"
	YouXiangCorp = "youxiang"

	SoyoungCorpNum  = 1
	YouXiangCorpNum = 2
)

const (
	// 朋友圈创建来源
	MomentCreateTypeCorp     = 0 // 企业
	MomentCreateTypePersonal = 1 // 个人

	// 成员发表状态。0:未发表 1：已发表
	MomentPublishStatusNo  = 0 // 未发表
	MomentPublishStatusYes = 1 // 已发表

	// 互动类型
	ActionTypeComment = "comment" // 评论
	ActionTypeLike    = "like"    // 点赞

	// 互动人类型
	ActionUserTypeUser         = 1 // 员工
	ActionUserTypeExternalUser = 2 // 客户
)

const (
	ChatMsgTypeText = "text" // 文本
)

// 企业微信加c二维码 自定义参数 业务类型定义
const (
	ContactWayBizTypeSelectedProductDetailGroup = "1_001"
	ContactWayBizTypeSelectedProductDetailC     = "1_002"
	ContactWayBizTypeDoctorMasterDetailC        = "1_003"
	ContactWayBizTypeAcnGroup                   = "1_004"
	ContactWayBizTypeAcnC                       = "1_005"
	ContactWayBizTypeXHSClear                   = "2_001"
	ContactWayBizTypeWeChatAddUser              = "3_001"
	ContactWayBizTypeChainOldBringNewC          = "4_001"
)

// 客户相关事件
const ChangeExternalContactEvent = "change_external_contact"
const (
	AddExternalContactChangeType     = "add_external_contact"      // 添加企业客户事件
	EditExternalContactChangeType    = "edit_external_contact"     // 编辑企业客户事件
	AddHalfExternalContactChangeType = "add_half_external_contact" // 外部联系人免验证添加成员事件
	DelExternalContactChangeType     = "del_external_contact"      // 删除企业客户事件，成员删除外部联系人时，回调该事件
	DelFollowUserChangeType          = "del_follow_user"           // 删除跟进成员事件，成员被外部联系人删除时，回调该事件
	TransferFailChangeType           = "transfer_fail"             // 客户接替失败事件，企业将客户分配给新的成员接替后，客户添加失败时回调该事件
)

// 员工相关事件
const StaffChangeContactEvent = "change_contact"
const (
	StaffCreateUser = "create_user" // 添加事件
	StaffUpdateUser = "update_user" // 编辑事件
	StaffDeleteUser = "delete_user" //
)

// 客户群相关事件
const ChangeExternalChatEvent = "change_external_chat"
const (
	ExternalChatCreateChangeType  = "create"  // 客户群创建事件，有新增客户群时，回调该事件
	ExternalChatUpdateChangeType  = "update"  // 客户群变更事件，客户群被修改后（群名变更，群成员增加或移除，群主变更，群公告变更），回调该事件
	ExternalChatDismissChangeType = "dismiss" // 客户群解散事件，当客户群被群主解散后，回调该事件
)
const (
	ExternalChatUpdateDetailAddMember    = "add_member"    // 成员入群
	ExternalChatUpdateDetailDelMember    = "del_member"    // 成员退群
	ExternalChatUpdateDetailChangeOwner  = "change_owner"  // 群主变更
	ExternalChatUpdateDetailChangeName   = "change_name"   // 群名变更
	ExternalChatUpdateDetailChangeNotice = "change_notice" // 群公告变更
)

// 企业客户标签相关事件
const ChangeExternalTagEvent = "change_external_tag"
const (
	ExternalTagCreateChangeType  = "create"  // 企业客户标签创建事件
	ExternalTagUpdateChangeType  = "update"  // 企业客户标签变更事件
	ExternalTagDeleteChangeType  = "delete"  // 企业客户标签删除事件
	ExternalTagShuffleChangeType = "shuffle" // 企业客户标签重排事件
)
const (
	TagTypeTag      = "tag"       // 类型-标签
	TagTypeTagGroup = "tag_group" // 类型-标签组
)

func GetContactWayQrCodeBizType() map[string]string {
	return map[string]string{
		ContactWayBizTypeSelectedProductDetailGroup: "新氧严选小程序商品详情页进群", // https://soyoung.feishu.cn/wiki/D8u9wg0pKiaVFyknj4TcINu9n8b
		ContactWayBizTypeSelectedProductDetailC:     "新氧严选小程序商品详情页加C", // https://soyoung.feishu.cn/wiki/D8u9wg0pKiaVFyknj4TcINu9n8b
		ContactWayBizTypeXHSClear:                   "小红书电话线索",        // https://soyoung.feishu.cn/wiki/Y7DZwI8iJip7OEkG19JcbffOnNd
		ContactWayBizTypeWeChatAddUser:              "企业微信添加员工",       // https://soyoung.feishu.cn/wiki/Upqow3U5BiWgZwks7lwcnOYAnQe
		ContactWayBizTypeDoctorMasterDetailC:        "万支大师",
		ContactWayBizTypeAcnGroup:                   "新氧严选小程序acn进群", // https://soyoung.feishu.cn/wiki/TaL9waMjhirwGSk2kzvc9InHnNe
		ContactWayBizTypeAcnC:                       "新氧严选小程序acn加c", // https://soyoung.feishu.cn/wiki/TaL9waMjhirwGSk2kzvc9InHnNe
	}
}
