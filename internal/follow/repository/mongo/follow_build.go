package mongo

import (
	"context"

	"github.com/hoag/go-social-feed/internal/follow/repository"
	"github.com/hoag/go-social-feed/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo impleRepository) buildModels(ctx context.Context, sc models.Scope, opts repository.CreateOptions) (models.Follow, error) {
	now := repo.clock()

	followerID, err := primitive.ObjectIDFromHex(opts.FollowerID)
	if err != nil {
		repo.l.Errorf(ctx, "reaction.repository.buildModels.ObjectIDFromHex: %v", err)
		return models.Follow{}, err
	}

	followeeID, err := primitive.ObjectIDFromHex(opts.FolloweeID)
	if err != nil {
		repo.l.Errorf(ctx, "reaction.repository.buildModels.ObjectIDFromHex: %v", err)
		return models.Follow{}, err
	}

	follow := models.Follow{
		ID:         repo.db.NewObjectID(),
		FollowerID: followerID,
		FolloweeID: followeeID,
		CreatedAt:  now,
	}

	return follow, nil
}
