package notification

type Notification struct {
	Content       string
	Heading       string
	UserIDs       []string
	CreatedUserID string
	En            MultiLangObj
	Ja            MultiLangObj
	Data          NotiData
	Source        SourceType
}

type MultiLangObj struct {
	Heading string
	Content string
}

type NotiData struct {
	Data     interface{}
	Activity ActivityType
}

type GetNotiHeadingInput struct {
	From SourceType
	Lang string
}

type GetNotiContentInput struct {
	From       SourceType
	Lang       string
	TaggerName string
	Content    string
}
