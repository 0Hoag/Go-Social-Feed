package notification

type ActivityType string

const (
	ActivityPostDetail ActivityType = "POST_DETAIL"
)

type SourceType string

const (
	SourceNewsFeed      SourceType = "news_feed"
	SourcePostCreate    SourceType = "post_created"
	SourcePostReactPost SourceType = "react_post"
	SourceTagUser       SourceType = "tag_user"
)
