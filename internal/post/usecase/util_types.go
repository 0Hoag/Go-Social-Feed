package usecase

import (
	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/internal/resource/notification"
)

type getPostNotiContent struct {
	P            models.Post
	Type         notification.SourceType
	TaggedTarget []string
	Content      string
	TaggerName   string
}
