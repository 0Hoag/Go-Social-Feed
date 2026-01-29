package mongo

import (
	"context"

	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/internal/users/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo impleRepository) buildModels(ctx context.Context, opts repository.CreateOptions) (models.User, error) {
	now := repo.clock()

	roles := make([]primitive.ObjectID, len(opts.Roles))
	for i, role := range opts.Roles {
		id, err := primitive.ObjectIDFromHex(role)
		if err != nil {
			repo.l.Errorf(ctx, "users.repo.mongo.user_build.ObjectIDFromHex: %v", err)
			return models.User{}, err
		}
		roles[i] = id
	}

	permissions := make([]primitive.ObjectID, len(opts.Permissions))
	for i, permission := range opts.Permissions {
		id, err := primitive.ObjectIDFromHex(permission)
		if err != nil {
			repo.l.Errorf(ctx, "users.repo.mongo.user_build.ObjectIDFromHex: %v", err)
			return models.User{}, err
		}

		permissions[i] = id
	}

	user := models.User{
		ID:          repo.db.NewObjectID(),
		Username:    opts.UserName,
		AvatarURL:   opts.AvatarURL,
		Phone:       opts.Phone,
		Password:    opts.Password,
		Birthday:    opts.Birthday,
		Roles:       roles,
		Permissions: permissions,
		CreatedAt:   now,
		UpdatedAt:   now,
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
