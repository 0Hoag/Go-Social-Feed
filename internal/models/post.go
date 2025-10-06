package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID           primitive.ObjectID   `bson:"_id,omitempty"`
	Pin          bool                 `bson:"pin"`
	Content      string               `bson:"content,omitempty"`
	FileIDs      []primitive.ObjectID `bson:"file_ids,omitempty"`
	TaggedTarget []primitive.ObjectID `bson:"tagged_target,omitempty"`
	Permission   PrivacyType          `bson:"permission,omitempty"`
	AuthorID     primitive.ObjectID   `bson:"author_id"`

	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty"`
}

type PrivacyType string

const (
	PrivacyTypePublic  PrivacyType = "public"
	PrivacyTypePrivate PrivacyType = "justme"
)
