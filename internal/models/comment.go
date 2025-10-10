package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	PostID      primitive.ObjectID `bson:"post_id"`
	AuthorID    primitive.ObjectID `bson:"author_id"`
	Content     string             `bson:"content"`
	Attachments []Attachment       `bson:"attachments,omitempty"`

	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty"`
}

type Attachment struct {
	Type string `bson:"type"` // "gif", "image", "video", etc.
	URL  string `bson:"url"`
}
