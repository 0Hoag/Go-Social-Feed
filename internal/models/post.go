package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID  primitive.ObjectID `bson:"author_id"`
	Content   string             `bson:"content,omitempty"`
	MediaURLs []string           `bson:"media_urls,omitempty"`

	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty"`
}
