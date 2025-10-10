package rabbitmq

import "github.com/hoag/go-social-feed/internal/resource/notification"

type DeletePostRelationMsg struct {
	PostID string `json:"post_id"`
	UserID string `json:"user_id"`
}

type NotiData struct {
	Data     interface{}               `json:"data"`
	Activity notification.ActivityType `json:"activity"`
}

type MultiLangObj struct {
	Heading string `json:"heading"`
	Content string `json:"content"`
}

type PublishNotiMsg struct {
	Content       string                  `json:"content"`
	Heading       string                  `json:"heading"`
	UserIDs       []string                `json:"user_ids"`
	CreatedUserID string                  `json:"created_user_id"`
	En            MultiLangObj            `json:"en"`
	Ja            MultiLangObj            `json:"ja"`
	Data          NotiData                `json:"data"`
	Source        notification.SourceType `json:"source"`
}
