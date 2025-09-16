package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Reaction represents a user's reaction to a post (like)
type Reaction struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	PostID    primitive.ObjectID `bson:"post_id"`
	UserID    primitive.ObjectID `bson:"user_id"`
	Type      Type               `bson:"type"` // e.g., "like"; can extend later
	CreatedAt time.Time          `bson:"created_at"`
}

type Type string

const (
	Like Type = "like"
	Love Type = "love"
	Haha Type = "haha"
)
