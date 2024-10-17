package config

import (
	"fmt"
)

const AppAccessTokenKey = "qywx_app_access_token_%s_%s"
const ContactAccessTokenKey = "qywx_contact_access_token_%s_%s"
const CustomerAccessTokenKey = "qywx_customer_access_token_%s_%s"

// 企微机器人
const (
	TaskTrackerCommandExecLock = "task:tracker:command:exec:lock:%s"
)

// 动态二维码透传回调参数
const AddContactWayCallbackKey = "add_contact_way_%s_%s"
const AddContactWayConfigKey = "add_contact_way_config"
const AddContactWayCallbackEx = 3600 * 24 * 2

func GetCacheKey(key string, args ...any) string {
	return fmt.Sprintf(key, args...)
}
