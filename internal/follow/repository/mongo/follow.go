package mongo

import (
	"context"
	"fmt"

	"github.com/hoag/go-social-feed/internal/follow/repository"
	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/pkg/mongo"
	"github.com/hoag/go-social-feed/pkg/paginator"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	followCollection = "follow"
)

func (repo impleRepository) getFollowCollection() mongo.Collection {
	collName := fmt.Sprintf("%s", followCollection)
	return repo.db.Collection(collName)
}

func (repo impleRepository) Create(ctx context.Context, sc models.Scope, opts repository.CreateOptions) (models.Follow, error) {
	col := repo.getFollowCollection()

	m, err := repo.buildModels(ctx, sc, opts)
	if err != nil {
		repo.l.Errorf(ctx, "follows.mongo.Create.buildModels: %v", err)
		return models.Follow{}, err
	}

	_, err = col.InsertOne(ctx, m)
	if err != nil {
		repo.l.Errorf(ctx, "follows.mogno.Create.InsertOne: %v", err)
		return models.Follow{}, err
	}

	return m, nil
}

func (repo impleRepository) Detail(ctx context.Context, sc models.Scope, id string) (models.Follow, error) {
	col := repo.getFollowCollection()

	filter, err := repo.buildDetailQuery(ctx, sc, id)
	if err != nil {
		repo.l.Errorf(ctx, "follows.mongo.Detail.buildDetailQuery: %v", err)
		return models.Follow{}, err
	}

	var m models.Follow
	err = col.FindOne(ctx, filter).Decode(&m)
	if err != nil {
		repo.l.Errorf(ctx, "follows.mongo.Detail.FindOne: %v", err)
		return models.Follow{}, err
	}

	return m, nil
}

func (repo impleRepository) List(ctx context.Context, sc models.Scope, opts repository.ListOptions) ([]models.Follow, error) {
	col := repo.getFollowCollection()

	filter, err := repo.buildListQuery(ctx, sc, opts)
	if err != nil {
		repo.l.Errorf(ctx, "follows.mongo.List.buildListQuery: %v", err)
		return []models.Follow{}, err
	}

	cur, err := col.Find(ctx, filter)
	if err != nil {
		repo.l.Errorf(ctx, "follows.mongo.List.buildListQuery: %v", err)
		return []models.Follow{}, err
	}

	var ms []models.Follow
	err = cur.All(ctx, ms)
	if err != nil {
		repo.l.Errorf(ctx, "follows.mongo.List.All: %v", err)
		return []models.Follow{}, err
	}

	return ms, nil
}

func (repo impleRepository) Get(ctx context.Context, sc models.Scope, opts repository.GetOptions) ([]models.Follow, paginator.Paginator, error) {
	col := repo.getFollowCollection()

	filter, err := repo.buildGetQuery(ctx, sc, opts)
	if err != nil {
		repo.l.Errorf(ctx, "follows.mongo.Get.buildGetQuery: %v", err)
		return []models.Follow{}, paginator.Paginator{}, err
	}

	cur, err := col.Find(ctx, filter, options.Find().
		SetLimit(opts.PagQuery.Limit).
		SetSkip(opts.PagQuery.Offset()))
	if err != nil {
		repo.l.Errorf(ctx, "follows.mongo.Get.Find: %v", err)
		return []models.Follow{}, paginator.Paginator{}, err
	}

	var ms []models.Follow
	err = cur.All(ctx, &ms)
	if err != nil {
		repo.l.Errorf(ctx, "follows.mongo.Get.All: %v", err)
		return []models.Follow{}, paginator.Paginator{}, err
	}

	total, err := col.CountDocuments(ctx, filter)
	if err != nil {
		repo.l.Errorf(ctx, "follows.mongo.Get.CountDocuments: %v", err)
		return []models.Follow{}, paginator.Paginator{}, err
	}

	return ms, paginator.Paginator{
		Total:       total,
		Count:       int64(len(ms)),
		PerPage:     opts.PagQuery.Limit,
		CurrentPage: opts.PagQuery.Page,
	}, nil
}

func (repo impleRepository) Delete(ctx context.Context, sc models.Scope, id string) error {
	col := repo.getFollowCollection()

	filter, err := repo.buildDetailQuery(ctx, sc, id)
	if err != nil {
		repo.l.Errorf(ctx, "follows.mongo.Delete.buildDetailQuery: %v", err)
		return err
	}

	_, err = col.DeleteOne(ctx, filter)
	if err != nil {
		repo.l.Errorf(ctx, "follows.mongo.Delete.DeleteOne: %v", err)
		return err
	}

	return nil
}
