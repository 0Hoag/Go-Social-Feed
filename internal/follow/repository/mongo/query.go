package mongo

import (
	"context"

	"github.com/hoag/go-social-feed/internal/follow/repository"
	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo impleRepository) buildDetailQuery(ctx context.Context, sc models.Scope, id string) (bson.M, error) {
	filter, err := mongo.BuildScopeQuery(ctx, repo.l, sc)
	if err != nil {
		repo.l.Errorf(ctx, "follow.mongo.buildDetailQuery.BuildScopeQuery: %v", err)
		return bson.M{}, err
	}

	filter = mongo.BuildQueryWithSoftDelete(filter)

	filter["_id"], err = primitive.ObjectIDFromHex(id)
	if err != nil {
		repo.l.Errorf(ctx, "follow.mongo.buildDetailQuery.BuildQueryWithSoftDelete: %v", err)
		return bson.M{}, err
	}

	return filter, nil
}

func (repo impleRepository) buildListQuery(ctx context.Context, sc models.Scope, opts repository.ListOptions) (bson.M, error) {
	filter, err := mongo.BuildScopeQuery(ctx, repo.l, sc)
	if err != nil {
		repo.l.Errorf(ctx, "follow.mongo.buildListQuery.BuildScopeQuery: %v", err)
		return bson.M{}, err
	}

	filter = mongo.BuildQueryWithSoftDelete(filter)

	if opts.ID != "" {
		filter["_id"], err = primitive.ObjectIDFromHex(opts.ID)
		if err != nil {
			repo.l.Errorf(ctx, "follow.mongo.buildListQuery.ObjectIDFromHex: %v", err)
			return bson.M{}, err
		}
	}

	mIDs := make([]primitive.ObjectID, len(opts.IDs))
	if len(opts.IDs) > 0 {
		for _, id := range opts.IDs {
			mID, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				repo.l.Errorf(ctx, "follow.mongo.buildListQuery.ObjectIDFromHex: %v", err)
				return bson.M{}, err
			}
			mIDs = append(mIDs, mID)
		}
		filter["_id"] = bson.M{"$in": mIDs}
	}

	if opts.Filter.AuthorID != "" {
		filter["author_id"], err = primitive.ObjectIDFromHex(opts.AuthorID)
		if err != nil {
			repo.l.Errorf(ctx, "follow.mongo.buildListQuery.ObjectIDFromHex: %v", err)
			return bson.M{}, err
		}
	}

	if opts.Filter.FolloweeID != "" {
		filter["followee_id"], err = primitive.ObjectIDFromHex(opts.FolloweeID)
		if err != nil {
			repo.l.Errorf(ctx, "follow.mongo.buildListQuery.ObjectIDFromHex: %v", err)
			return bson.M{}, err
		}
	}

	return filter, nil
}

func (repo impleRepository) buildGetQuery(ctx context.Context, sc models.Scope, opts repository.GetOptions) (bson.M, error) {
	filter, err := mongo.BuildScopeQuery(ctx, repo.l, sc)
	if err != nil {
		repo.l.Errorf(ctx, "follow.mongo.buildGetQuery.BuildScopeQuery: %v", err)
		return bson.M{}, err
	}

	filter = mongo.BuildQueryWithSoftDelete(filter)

	if opts.ID != "" {
		filter["_id"], err = primitive.ObjectIDFromHex(opts.ID)
		if err != nil {
			repo.l.Errorf(ctx, "follow.mongo.buildGetQuery.ObjectIDFromHex: %v", err)
			return bson.M{}, err
		}
	}

	mIDs := make([]primitive.ObjectID, len(opts.IDs))
	if len(opts.IDs) > 0 {
		for _, id := range opts.IDs {
			mID, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				repo.l.Errorf(ctx, "follow.mongo.buildGetQuery.ObjectIDFromHex: %v", err)
				return bson.M{}, err
			}
			mIDs = append(mIDs, mID)
		}
		filter["_id"] = bson.M{"$in": mIDs}
	}

	if opts.Filter.AuthorID != "" {
		filter["author_id"], err = primitive.ObjectIDFromHex(opts.AuthorID)
		if err != nil {
			repo.l.Errorf(ctx, "follow.mongo.buildGetQuery.ObjectIDFromHex: %v", err)
			return bson.M{}, err
		}
	}

	if opts.Filter.FolloweeID != "" {
		filter["followee_id"], err = primitive.ObjectIDFromHex(opts.FolloweeID)
		if err != nil {
			repo.l.Errorf(ctx, "follow.mongo.buildGetQuery.ObjectIDFromHex: %v", err)
			return bson.M{}, err
		}
	}

	return filter, nil
}
