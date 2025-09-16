package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Follow represents follow relationship (follower -> followee)
type Follow struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	FollowerID primitive.ObjectID `bson:"follower_id"`
	FolloweeID primitive.ObjectID `bson:"followee_id"`

	CreatedAt time.Time `bson:"created_at"`
}
