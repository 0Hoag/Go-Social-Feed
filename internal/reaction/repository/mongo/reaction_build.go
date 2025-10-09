package mongo

import (
	"context"

	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/internal/reaction/repository"
	"github.com/hoag/go-social-feed/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo impleRepository) buildModels(ctx context.Context, sc models.Scope, opts repository.CreateOptions) (models.Reaction, error) {
	now := repo.clock()

	postID, err := primitive.ObjectIDFromHex(opts.PostID)
	if err != nil {
		repo.l.Errorf(ctx, "reaction.repository.buildModels.ObjectIDFromHex: %v", err)
		return models.Reaction{}, err
	}

	reaction := models.Reaction{
		ID:        repo.db.NewObjectID(),
		PostID:    postID,
		UserID:    mongo.ObjectIDFromHexOrNil(sc.UserID),
		CreatedAt: now,
	}

	return reaction, nil
}
