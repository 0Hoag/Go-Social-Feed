package mongo

import (
	"context"

	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/internal/post/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo impleRepository) buildModels(ctx context.Context, sc models.Scope, opts repository.CreateOptions) (bson.M, error) {
	now := repo.clock()

	tmpTagged := models.TaggedTarget{}
	fileIDs := make([]primitive.ObjectID, len(opts.FileIDs))

	if len(opts.TaggedTarget) > 0 {
		userIDs := make([]primitive.ObjectID, len(opts.TaggedTarget))
		for i, id := range opts.TaggedTarget {
			objID, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				repo.l.Errorf(ctx, "post.repository.mongo.post_build.buildModels.primitive.ObjectIDFromHex: %v", err)
				return bson.M{}, err
			}
			userIDs[i] = objID
		}
		tmpTagged = userIDs
	}

	if len(opts.FileIDs) > 0 {
		for i, id := range opts.FileIDs {
			objID, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				repo.l.Errorf(ctx, "post.repository.mongo.post_build.buildModels.primitive.ObjectIDFromHex: %v", err)
				return bson.M{}, err
			}
			fileIDs[i] = objID
		}
	}

}
