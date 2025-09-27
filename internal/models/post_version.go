package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostVersion struct {
	ID         primitive.ObjectID   `bson:"_id"`
	PostID     primitive.ObjectID   `bson:"post_id"`
	Version    int                  `bson:"version"`
	Content    string               `bson:"content,omitempty"`
	FileIDs    []primitive.ObjectID `bson:"file_ids,omitempty"`
	Permission PrivacyType          `bson:"permission,omitempty"`
	CreatedAt  time.Time            `bson:"created_at"`
	UpdatedAt  time.Time            `bson:"updated_at"`
	DeletedAt  *time.Time           `bson:"deleted_at,omitempty"`
}

type PrivacyType string

const (
	PrivacyTypePublic  PrivacyType = "public"
	PrivacyTypePrivate PrivacyType = "justme"
)
