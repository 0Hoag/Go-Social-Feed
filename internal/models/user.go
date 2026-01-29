package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty"`
	Username    string               `bson:"username"`
	Phone       string               `bson:"phone,omitempty"`
	Password    string               `bson:"password,omitempty"`
	AvatarURL   string               `bson:"avatar_url,omitempty"`
	Bio         string               `bson:"bio,omitempty"`
	Birthday    time.Time            `bson:"birthday,omitempty"`
	Roles       []primitive.ObjectID `bson:"roles,omitempty"`
	Permissions []primitive.ObjectID `bson:"permissions,omitempty"`
	CreatedAt   time.Time            `bson:"created_at"`
	UpdatedAt   time.Time            `bson:"updated_at"`
	DeletedAt   *time.Time           `bson:"deleted_at,omitempty"`
}

type Role string
type Permission string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

const (
	PermissionCreatePost Permission = "create_post"
	PermissionUpdatePost Permission = "update_post"
	PermissionDeletePost Permission = "delete_post"
	PermissionReadPost   Permission = "read_post"
)

type Roles struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        Role               `bson:"name"`
	Permissions []string           `bson:"permissions"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	DeletedAt   *time.Time         `bson:"deleted_at,omitempty"`
}

type Permissions struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      Permission         `bson:"name"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	DeletedAt *time.Time         `bson:"deleted_at,omitempty"`
}
