package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Comment represents a comment under a post
type Comment struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	PostID   primitive.ObjectID `bson:"post_id"`
	AuthorID primitive.ObjectID `bson:"author_id"`
	Content  string             `bson:"content"`

	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty"`
}
