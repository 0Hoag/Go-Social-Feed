package mongo

import (
	"context"

	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/internal/post/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo impleRepository) buildModels(ctx context.Context, sc models.Scope, opts repository.CreateOptions) (models.Post, error) {
	now := repo.clock()

	tmpTagged := make([]primitive.ObjectID, len(opts.TaggedTarget))
	fileIDs := make([]primitive.ObjectID, len(opts.FileIDs))

	if len(opts.TaggedTarget) > 0 {
		for i, id := range opts.TaggedTarget {
			objID, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				repo.l.Errorf(ctx, "post.mongo.post_build.buildModels.ObjectIDFromHex: %v", err)
				return models.Post{}, err
			}
			tmpTagged[i] = objID
		}
	}

	if len(opts.FileIDs) > 0 {
		for i, id := range opts.FileIDs {
			objID, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				repo.l.Errorf(ctx, "post.mongo.post_build.buildModels.ObjectIDFromHex: %v", err)
				return models.Post{}, err
			}
			fileIDs[i] = objID
		}
	}

	authorID, err := primitive.ObjectIDFromHex(sc.UserID)
	if err != nil {
		repo.l.Debugf(ctx, "post.mongo.post_build.buildModels.ObjectIDFromHex: %v", err)
		return models.Post{}, err
	}

	post := models.Post{
		ID:           repo.db.NewObjectID(),
		Pin:          opts.Pin,
		Content:      opts.Content,
		FileIDs:      fileIDs,
		TaggedTarget: tmpTagged,
		Permission:   models.PrivacyType(opts.Permission),
		AuthorID:     authorID,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	return post, nil
}

func (repo impleRepository) buildUpdateModels(ctx context.Context, sc models.Scope, opts repository.UpdateOptions) (bson.M, error) {
	now := repo.clock()

	set := bson.M{}

	fileIDs := make([]primitive.ObjectID, len(opts.FileIDs))
	taggedIDs := make([]primitive.ObjectID, len(opts.TaggedTarget))

	if len(opts.FileIDs) > 0 {
		for _, id := range opts.FileIDs {
			fID, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				repo.l.Errorf(ctx, "post.mongo.buildUpdateModels.ObjectIDFromHex: %v", err)
				return bson.M{}, err
			}
			fileIDs = append(fileIDs, fID)
		}
	}

	if len(opts.TaggedTarget) > 0 {
		for _, id := range opts.TaggedTarget {
			tID, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				repo.l.Errorf(ctx, "post.mongo.buildUpdateModels.ObjectIDFromHex: %v", err)
				return bson.M{}, err
			}
			taggedIDs = append(fileIDs, tID)
		}
	}

	if opts.Content != "" {
		set["content"] = opts.Content
	}

	if len(fileIDs) > 0 {
		set["file_ids"] = fileIDs
	}

	if len(taggedIDs) > 0 {
		set["tagged_target"] = taggedIDs
	}

	if opts.Pin {
		set["pin"] = opts.Pin
	}

	set["updated_at"] = now

	return set, nil
}
