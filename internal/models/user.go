package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Username     string             `bson:"username"`
	Phone        string             `bson:"phone,omitempty"`
	PasswordHash string             `bson:"password_hash,omitempty"`
	AvatarURL    string             `bson:"avatar_url,omitempty"`
	Bio          string             `bson:"bio,omitempty"`
	Birthday     time.Time          `bson:"birthday,omitempty"`

	Roles       []string `bson:"roles,omitempty"`
	Permissions []string `bson:"permissions,omitempty"`

	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty"`
}
