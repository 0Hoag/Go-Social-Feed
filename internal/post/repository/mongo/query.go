package mongo

import (
	"context"

	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/internal/post/repository"
	"github.com/hoag/go-social-feed/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo impleRepository) buildDetailQuery(ctx context.Context, sc models.Scope, id string) (bson.M, error) {
	filter := bson.M{}

	filter = mongo.BuildQueryWithSoftDelete(filter)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		repo.l.Errorf(ctx, "post.mongo.buildDetailQuery.ObjectIDFromHex: %v", err)
		return bson.M{}, err
	}

	filter["_id"] = objectID

	return filter, nil
}

func (repo impleRepository) buildListQuery(ctx context.Context, sc models.Scope, opts repository.ListOptions) (bson.M, error) {
	filter := bson.M{}

	filter = mongo.BuildQueryWithSoftDelete(filter)

	if opts.ID != "" {
		objectID, err := primitive.ObjectIDFromHex(opts.ID)
		if err != nil {
			repo.l.Errorf(ctx, "post.mongo.buildListQuery.ObjectIDFromHex: %v", err)
			return bson.M{}, err
		}
		filter["_id"] = objectID
	}

	mIDs := make([]primitive.ObjectID, len(opts.IDs))
	if len(opts.IDs) > 0 {
		for _, id := range opts.IDs {
			mID, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				repo.l.Errorf(ctx, "post.mongo.buildListQuery.ObjectIDFromHex: %v", err)
				return bson.M{}, err
			}
			mIDs = append(mIDs, mID)
		}
		filter["_id"] = bson.M{"$in": mIDs}
	}

	if opts.AuthorID != "" {
		objectID, err := primitive.ObjectIDFromHex(opts.AuthorID)
		if err != nil {
			repo.l.Errorf(ctx, "post.mongo.buildListQuery.ObjectIDFromHex: %v", err)
			return bson.M{}, err
		}
		filter["author_id"] = objectID
	}

	filter["pin"] = opts.Pin

	return filter, nil
}

func (repo impleRepository) buildGetQuery(ctx context.Context, sc models.Scope, opts repository.GetOptions) (bson.M, error) {
	filter := bson.M{}

	filter = mongo.BuildQueryWithSoftDelete(filter)

	if opts.ID != "" {
		objectID, err := primitive.ObjectIDFromHex(opts.ID)
		if err != nil {
			repo.l.Errorf(ctx, "post.mongo.buildListQuery.ObjectIDFromHex: %v", err)
			return bson.M{}, err
		}
		filter["_id"] = objectID
	}

	mIDs := make([]primitive.ObjectID, len(opts.IDs))
	if len(opts.IDs) > 0 {
		for _, id := range opts.IDs {
			mID, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				repo.l.Errorf(ctx, "post.mongo.buildListQuery.ObjectIDFromHex: %v", err)
				return bson.M{}, err
			}
			mIDs = append(mIDs, mID)
		}
		filter["_id"] = bson.M{"$in": mIDs}
	}

	if opts.AuthorID != "" {
		objectID, err := primitive.ObjectIDFromHex(opts.AuthorID)
		if err != nil {
			repo.l.Errorf(ctx, "post.mongo.buildListQuery.ObjectIDFromHex: %v", err)
			return bson.M{}, err
		}
		filter["author_id"] = objectID
	}

	filter["pin"] = opts.Pin

	return filter, nil
}

func (repo impleRepository) buildGetOneQuery(ctx context.Context, sc models.Scope, opts repository.GetOneOptions) (bson.M, error) {
	filter := bson.M{}

	filter = mongo.BuildQueryWithSoftDelete(filter)

	if opts.ID != "" {
		objectID, err := primitive.ObjectIDFromHex(opts.ID)
		if err != nil {
			repo.l.Errorf(ctx, "post.mongo.buildGetOneQuery.ObjectIDFromHex: %v", err)
			return bson.M{}, err
		}
		filter["_id"] = objectID
		if err != nil {
			repo.l.Errorf(ctx, "post.mongo.buildGetOneQuery.ObjectIDFromHex: %v", err)
			return bson.M{}, err
		}
	}

	if opts.AuthorID != "" {
		objectID, err := primitive.ObjectIDFromHex(opts.AuthorID)
		if err != nil {
			repo.l.Errorf(ctx, "post.mongo.buildGetOneQuery.ObjectIDFromHex: %v", err)
			return bson.M{}, err
		}
		filter["author_id"] = objectID
	}

	if opts.SourceURL != "" {
		filter["source_url"] = opts.SourceURL
	}

	return filter, nil
}
