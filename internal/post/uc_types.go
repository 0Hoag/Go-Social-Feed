package post

import (
	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/internal/resource/notification"
	"github.com/hoag/go-social-feed/pkg/paginator"
)

// Post
type CreateInput struct {
	Pin          bool
	Content      string
	FileIDs      []string
	TaggedTarget []string
	Permission   string
}

type Filter struct {
	ID       string
	IDs      []string
	Pin      bool
	AuthorID string
}

type ListInput struct {
	Filter
}

type GetInput struct {
	Filter
	PagQuery paginator.PaginatorQuery
}

type GetOutput struct {
	Posts     []models.Post
	Paginator paginator.Paginator
}

type UpdateInput struct {
	ID           string
	Pin          bool
	Content      string
	FileIDs      []string
	TaggedTarget []string
	Permission   string
}

// Notification
type PublishNotiPostInput struct {
	PostID     string                  `json:"post_id"`
	ReceiverID string                  `json:"receiver_id,omitempty"`
	Type       notification.SourceType `json:"type"`
}

type NotificationInput struct {
	Post         models.Post
	TaggedTarget []string
}

type NotificationOutput struct {
	Users       []models.User
	SessionUser models.User
}

// Message
type DeleteCommentMsgInput struct {
	PostID string
}

type DeleteReactionMsgInput struct {
	PostID string
}
