package mongo

import (
	"context"

	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/internal/users/repository"
	"go.mongodb.org/mongo-driver/bson"
)

func (repo impleRepository) buildModels(ctx context.Context, sc models.Scope, opts repository.CreateOptions) (models.User, error) {
	now := repo.clock()

	user := models.User{
		ID:           repo.db.NewObjectID(),
		Username:     opts.UserName,
		Phone:        opts.Phone,
		PasswordHash: opts.PasswordHash,
		Birthday:     opts.Birthday,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	return user, nil
}

func (repo impleRepository) buildUpdateModels(ctx context.Context, sc models.Scope, opts repository.UpdateOptions) (bson.M, error) {
	now := repo.clock()

	set := bson.M{}

	if opts.AvatarURL != "" {
		set["avatar_url"] = opts.AvatarURL
	}

	if opts.UserName != "" {
		set["username"] = opts.UserName
	}

	set["updated_at"] = now

	return set, nil
}
