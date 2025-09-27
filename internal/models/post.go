package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Pin          bool               `bson:"pin"`
	TaggedTarget []string           `bson:"tagged_target,omitempty"`
	AuthorID     primitive.ObjectID `bson:"author_id"`

	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty"`
}
