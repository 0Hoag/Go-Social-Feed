package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reaction struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	PostID primitive.ObjectID `bson:"post_id"`
	UserID primitive.ObjectID `bson:"user_id"`
	Type   string             `bson:"type"`

	CreatedAt time.Time `bson:"created_at"`
}
