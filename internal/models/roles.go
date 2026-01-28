package models

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type Roles struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Name Role `bson:"name"`
	Permissions []string `bson:"permissions"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty"`
}

