package repository

import "github.com/hoag/go-social-feed/pkg/paginator"

// Post
type CreateOptions struct {
	Content      string
	FileIDs      []string
	Pin          bool
	TaggedTarget []string
}

type Filter struct {
	ID       string
	IDs      []string
	Pin      bool
	AuthorID string
}

type ListOptions struct {
	Filter
}

type GetOptions struct {
	Filter
	PagQuery paginator.PaginatorQuery
}

type UpdateOptions struct {
	ID           string
	Content      string
	FileIDs      []string
	Pin          bool
	TaggedTarget []string
}
