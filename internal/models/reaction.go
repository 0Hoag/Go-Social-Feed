package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reaction struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	PostID   primitive.ObjectID `bson:"post_id"`
	AuthorID primitive.ObjectID `bson:"author_id"`
	Type     ReactionType       `bson:"type"`

	CreatedAt time.Time `bson:"created_at"`
}

type ReactionType string

const (
	LikeReaction ReactionType = "like"
	LoveReaction ReactionType = "love"
)
