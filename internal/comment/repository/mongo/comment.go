package mongo

import (
	"context"

	"github.com/hoag/go-social-feed/internal/comment/repository"
	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/pkg/mongo"
	"github.com/hoag/go-social-feed/pkg/paginator"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	commentCollection = "comment"
)

func (repo impleRepository) getCommentCollection() mongo.Collection {
	return repo.db.Collection(commentCollection)
}

func (repo impleRepository) Create(ctx context.Context, sc models.Scope, opts repository.CreateOptions) (models.Comment, error) {
	col := repo.getCommentCollection()

	m, err := repo.buildModels(ctx, sc, opts)
	if err != nil {
		repo.l.Errorf(ctx, "comments.mongo.Create.buildModels: %v", err)
		return models.Comment{}, err
	}

	_, err = col.InsertOne(ctx, m)
	if err != nil {
		repo.l.Errorf(ctx, "comments.mogno.Create.InsertOne: %v", err)
		return models.Comment{}, err
	}

	return m, nil
}

func (repo impleRepository) Detail(ctx context.Context, sc models.Scope, id string) (models.Comment, error) {
	col := repo.getCommentCollection()

	filter, err := repo.buildDetailQuery(ctx, sc, id)
	if err != nil {
		repo.l.Errorf(ctx, "comments.mongo.Detail.buildDetailQuery: %v", err)
		return models.Comment{}, err
	}

	var m models.Comment
	err = col.FindOne(ctx, filter).Decode(&m)
	if err != nil {
		repo.l.Errorf(ctx, "comments.mongo.Detail.FindOne: %v", err)
		return models.Comment{}, err
	}

	return m, nil
}

func (repo impleRepository) GetOne(ctx context.Context, sc models.Scope, opts repository.GetOneOptions) (models.Comment, error) {
	col := repo.getCommentCollection()

	filter, err := repo.buildGetOneQuery(ctx, sc, opts.Filter)
	if err != nil {
		repo.l.Errorf(ctx, "comments.mongo.Detail.buildDetailQuery: %v", err)
		return models.Comment{}, err
	}

	var m models.Comment
	err = col.FindOne(ctx, filter).Decode(&m)
	if err != nil {
		repo.l.Errorf(ctx, "comments.mongo.Detail.FindOne: %v", err)
		return models.Comment{}, err
	}

	return m, nil
}

func (repo impleRepository) List(ctx context.Context, sc models.Scope, opts repository.ListOptions) ([]models.Comment, error) {
	col := repo.getCommentCollection()

	filter, err := repo.buildListQuery(ctx, sc, opts)
	if err != nil {
		repo.l.Errorf(ctx, "comments.mongo.List.buildListQuery: %v", err)
		return []models.Comment{}, err
	}

	cur, err := col.Find(ctx, filter)
	if err != nil {
		repo.l.Errorf(ctx, "comments.mongo.List.buildListQuery: %v", err)
		return []models.Comment{}, err
	}

	var ms []models.Comment
	err = cur.All(ctx, ms)
	if err != nil {
		repo.l.Errorf(ctx, "comments.mongo.List.All: %v", err)
		return []models.Comment{}, err
	}

	return ms, nil
}

func (repo impleRepository) Get(ctx context.Context, sc models.Scope, opts repository.GetOptions) ([]models.Comment, paginator.Paginator, error) {
	col := repo.getCommentCollection()

	filter, err := repo.buildGetQuery(ctx, sc, opts)
	if err != nil {
		repo.l.Errorf(ctx, "comments.mongo.Get.buildGetQuery: %v", err)
		return []models.Comment{}, paginator.Paginator{}, err
	}

	cur, err := col.Find(ctx, filter, options.Find().
		SetLimit(opts.PagQuery.Limit).
		SetSkip(opts.PagQuery.Offset()))
	if err != nil {
		repo.l.Errorf(ctx, "comments.mongo.Get.Find: %v", err)
		return []models.Comment{}, paginator.Paginator{}, err
	}

	var ms []models.Comment
	err = cur.All(ctx, &ms)
	if err != nil {
		repo.l.Errorf(ctx, "comments.mongo.Get.All: %v", err)
		return []models.Comment{}, paginator.Paginator{}, err
	}

	total, err := col.CountDocuments(ctx, filter)
	if err != nil {
		repo.l.Errorf(ctx, "comments.mongo.Get.CountDocuments: %v", err)
		return []models.Comment{}, paginator.Paginator{}, err
	}

	return ms, paginator.Paginator{
		Total:       total,
		Count:       int64(len(ms)),
		PerPage:     opts.PagQuery.Limit,
		CurrentPage: opts.PagQuery.Page,
	}, nil
}

func (repo impleRepository) Update(ctx context.Context, sc models.Scope, opts repository.UpdateOptions) (models.Comment, error) {
	col := repo.getCommentCollection()

	m, filter, err := repo.buildUpdateModels(ctx, sc, opts)
	if err != nil {
		repo.l.Errorf(ctx, "comments.mongo.Update.buildUpdateModels: %v", err)
		return models.Comment{}, nil
	}

	_, err = col.UpdateOne(ctx, filter, &m)
	if err != nil {
		repo.l.Errorf(ctx, "comments.mongo.Update.UpdateOne: %v", err)
		return models.Comment{}, nil
	}

	return m, nil
}

func (repo impleRepository) Delete(ctx context.Context, sc models.Scope, id string) error {
	col := repo.getCommentCollection()

	filter, err := repo.buildDetailQuery(ctx, sc, id)
	if err != nil {
		repo.l.Errorf(ctx, "comments.mongo.Delete.buildDetailQuery: %v", err)
		return err
	}

	_, err = col.DeleteOne(ctx, filter)
	if err != nil {
		repo.l.Errorf(ctx, "comments.mongo.Delete.DeleteOne: %v", err)
		return err
	}

	return nil
}
