package models

type Permission string

const (
	PermissionCreatePost Permission = "create_post"
	PermissionUpdatePost Permission = "update_post"
	PermissionDeletePost Permission = "delete_post"
	PermissionReadPost   Permission = "read_post"
)

type Permissions struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Name Permission `bson:"name"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty"`
}