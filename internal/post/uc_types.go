package post

import (
	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/internal/resource/notification"
	"github.com/hoag/go-social-feed/pkg/paginator"
)

// Post
type CreateInput struct {
	Pin          bool
	Title        string
	Content      string
	FileIDs      []string
	TaggedTarget []string
	Permission   string
	SourceURL    string
	FullContent  string
}

type Filter struct {
	ID        string
	IDs       []string
	Pin       bool
	AuthorID  string
	SourceURL string
}

type GetOneInput struct {
	Filter
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

// reaction
type CreateReactionInput struct {
	PostID string
	Type   models.ReactionType
}

type FilterReaction struct {
	ID     string
	IDs    []string
	UserID string
	Type   models.ReactionType
}

type ListReactionInput struct {
	FilterReaction
}

type GetReactionInput struct {
	FilterReaction
	PagQuery paginator.PaginatorQuery
}

type GetReactionOutput struct {
	Reactions []models.Reaction
	Paginator paginator.Paginator
}

// Comment
type CreateCommentInput struct {
	PostID  string
	Content string
	Attach  []models.Attachment
}

type FilterComment struct {
	ID     string
	IDs    []string
	PostID string
}

type ListCommentInput struct {
	FilterComment
}

type GetCommentInput struct {
	FilterComment
	PagQuery paginator.PaginatorQuery
}

type GetCommentOutput struct {
	Comments  []models.Comment
	Paginator paginator.Paginator
}

type UpdateCommentInput struct {
	PostID  string
	Content string
	Attach  []models.Attachment
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
