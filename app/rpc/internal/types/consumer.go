package types

import "time"

type TopicMessage struct {
	Topic     string `json:"topic"`
	From      string `json:"from"`
	Timestamp int64  `json:"timestamp"`
	Logid     string `json:"logid"`
}

type ExternalData struct {
	ToUserName     string    `json:"ToUserName"`
	FromUserName   string    `json:"FromUserName"`
	CreateTime     time.Time `json:"CreateTime"`
	MsgType        string    `json:"MsgType"`
	Event          string    `json:"Event"`
	ChangeType     string    `json:"ChangeType"`
	UserID         string    `json:"UserID"`
	ExternalUserID string    `json:"ExternalUserID"`
}
