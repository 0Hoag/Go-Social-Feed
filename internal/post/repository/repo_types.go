package repository

import (
	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/pkg/paginator"
)

// Post
type CreateOptions struct {
	Pin           bool
	Title         string
	TitleEn       string
	Content       string
	FullContent   string
	FullContentEn string
	FileIDs       []string
	TaggedTarget  []string
	Permission    string
	SourceURL     string
}

type Filter struct {
	ID        string
	IDs       []string
	Pin       bool
	AuthorID  string
	SourceURL string
}

type GetOneOptions struct {
	Filter
}

type ListOptions struct {
	Filter
}

type GetOptions struct {
	Filter
	PagQuery paginator.PaginatorQuery
}

type UpdateOptions struct {
	Post         models.Post
	Pin          bool
	Content      string
	FileIDs      []string
	TaggedTarget []string
	Permission   string
}

// Reaction
type CreateReactionOptions struct {
	PostID string
	Type   models.ReactionType
}

type FilterReaction struct {
	ID     string
	IDs    []string
	UserID string
	Type   models.ReactionType
}

type ListReactionOptions struct {
	FilterReaction
}

type GetReactionOptions struct {
	FilterReaction
	PagQuery paginator.PaginatorQuery
}

// Comment
type CreateCommentOptions struct {
	PostID  string
	Content string
	Attach  []models.Attachment
}

type GetOneCommentOptions struct {
	FilterComment
}

type FilterComment struct {
	ID     string
	IDs    []string
	PostID string
}

type ListCommentOptions struct {
	FilterComment
}

type GetCommentOptions struct {
	FilterComment
	PagQuery paginator.PaginatorQuery
}

type UpdateCommentOptions struct {
	Comment models.Comment
	Content string
	Attach  []models.Attachment
}
