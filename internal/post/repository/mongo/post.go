package mongo

import (
	"context"
	"fmt"

	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/internal/post/repository"
	"github.com/hoag/go-social-feed/pkg/mongo"
	"github.com/hoag/go-social-feed/pkg/paginator"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	postCollection = "post"
)

func (repo impleRepository) getPostCollection() mongo.Collection {
	collName := fmt.Sprintf("%s", postCollection)
	return repo.db.Collection(collName)
}

func (repo impleRepository) Create(ctx context.Context, sc models.Scope, opts repository.CreateOptions) (models.Post, error) {
	col := repo.getPostCollection()

	m, err := repo.buildModels(ctx, sc, opts)
	if err != nil {
		repo.l.Errorf(ctx, "post.mongo.Create.buildModels: %v", err)
		return models.Post{}, err
	}

	_, err = col.InsertOne(ctx, m)
	if err != nil {
		repo.l.Errorf(ctx, "post.mogno.Create.InsertOne: %v", err)
		return models.Post{}, err
	}

	return m, nil
}

func (repo impleRepository) Detail(ctx context.Context, sc models.Scope, id string) (models.Post, error) {
	col := repo.getPostCollection()

	filter, err := repo.buildDetailQuery(ctx, sc, id)
	if err != nil {
		repo.l.Errorf(ctx, "post.mongo.Detail.buildDetailQuery: %v", err)
		return models.Post{}, err
	}

	var m models.Post
	err = col.FindOne(ctx, filter).Decode(&m)
	if err != nil {
		repo.l.Errorf(ctx, "post.mongo.Detail.FindOne: %v", err)
		return models.Post{}, err
	}

	return m, nil
}

func (repo impleRepository) List(ctx context.Context, sc models.Scope, opts repository.ListOptions) ([]models.Post, error) {
	col := repo.getPostCollection()

	filter, err := repo.buildListQuery(ctx, sc, opts)
	if err != nil {
		repo.l.Errorf(ctx, "post.mongo.List.buildListQuery: %v", err)
		return []models.Post{}, err
	}

	cur, err := col.Find(ctx, filter)
	if err != nil {
		repo.l.Errorf(ctx, "post.mongo.List.buildListQuery: %v", err)
		return []models.Post{}, err
	}

	var ms []models.Post
	err = cur.All(ctx, ms)
	if err != nil {
		repo.l.Errorf(ctx, "post.mongo.List.All: %v", err)
		return []models.Post{}, err
	}

	return ms, nil
}

func (repo impleRepository) Get(ctx context.Context, sc models.Scope, opts repository.GetOptions) ([]models.Post, paginator.Paginator, error) {
	col := repo.getPostCollection()

	filter, err := repo.buildGetQuery(ctx, sc, opts)
	if err != nil {
		repo.l.Errorf(ctx, "post.mongo.Get.buildGetQuery: %v", err)
		return []models.Post{}, paginator.Paginator{}, err
	}

	cur, err := col.Find(ctx, filter, options.Find().
		SetLimit(opts.PagQuery.Limit).
		SetSkip(opts.PagQuery.Offset()))
	if err != nil {
		repo.l.Errorf(ctx, "post.mongo.Get.Find: %v", err)
		return []models.Post{}, paginator.Paginator{}, err
	}

	var ms []models.Post
	err = cur.All(ctx, ms)
	if err != nil {
		repo.l.Errorf(ctx, "post.mongo.Get.All: %v", err)
		return []models.Post{}, paginator.Paginator{}, err
	}

	total, err := col.CountDocuments(ctx, filter)
	if err != nil {
		repo.l.Errorf(ctx, "post.mongo.Get.CountDocuments: %v", err)
		return []models.Post{}, paginator.Paginator{}, err
	}

	return ms, paginator.Paginator{
		Total:       total,
		Count:       int64(len(ms)),
		PerPage:     opts.PagQuery.Limit,
		CurrentPage: opts.PagQuery.Page,
	}, nil
}
