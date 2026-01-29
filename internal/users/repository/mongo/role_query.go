package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo impleRepository) buildDetailRoleQuery(ctx context.Context, id string) (bson.M, error) {
	filter := bson.M{}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		repo.l.Errorf(ctx, "users.mongo.buildDetailRoleQuery.ObjectIDFromHex: %v", err)
		return bson.M{}, err
	}

	filter["_id"] = objectID

	return filter, nil
}
