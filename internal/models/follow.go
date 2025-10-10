package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Follow struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID   primitive.ObjectID `bson:"author_id"`
	FolloweeID primitive.ObjectID `bson:"followee_id"`

	CreatedAt time.Time  `bson:"created_at"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty"`
}
