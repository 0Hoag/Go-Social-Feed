package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID            primitive.ObjectID   `bson:"_id,omitempty"`
	Pin           bool                 `bson:"pin"`
	Title         string               `bson:"title,omitempty"`
	TitleEn       string               `bson:"title_en,omitempty"` // English Title
	Content       string               `bson:"content,omitempty"`
	FileIDs       []primitive.ObjectID `bson:"file_ids,omitempty"`
	TaggedTarget  []primitive.ObjectID `bson:"tagged_target,omitempty"`
	Permission    PrivacyType          `bson:"permission,omitempty"`
	AuthorID      primitive.ObjectID   `bson:"author_id"`
	SourceURL     string               `bson:"source_url,omitempty"`
	FullContent   string               `bson:"full_content,omitempty"`    // Complete article text (VI)
	FullContentEn string               `bson:"full_content_en,omitempty"` // Complete article text (EN)

	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty"`
}

type PrivacyType string

const (
	PrivacyTypePublic  PrivacyType = "public"
	PrivacyTypePrivate PrivacyType = "justme"
)
