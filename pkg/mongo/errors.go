package mongo

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrNoDocuments     = mongo.ErrNoDocuments
	ErrInvalidObjectID = errors.New("invalid object id")
)
