package mongo

import (
	"context"

	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/internal/post/repository"
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

func (repo impleRepository) CreateComment(ctx context.Context, sc models.Scope, opts repository.CreateCommentOptions) (models.Comment, error) {
	col := repo.getCommentCollection()

	m, err := repo.buildCommentModels(ctx, sc, opts)
	if err != nil {
		repo.l.Errorf(ctx, "post.mongo.CreateComment.buildModels: %v", err)
		return models.Comment{}, err
	}

	_, err = col.InsertOne(ctx, m)
	if err != nil {
		repo.l.Errorf(ctx, "post.mogno.CreateComment.InsertOne: %v", err)
		return models.Comment{}, err
	}

	return m, nil
}

func (repo impleRepository) DetailComment(ctx context.Context, sc models.Scope, id string) (models.Comment, error) {
	col := repo.getCommentCollection()

	filter, err := repo.buildDetailCommentQuery(ctx, sc, id)
	if err != nil {
		repo.l.Errorf(ctx, "post.mongo.DetailComment.buildDetailCommentQuery: %v", err)
		return models.Comment{}, err
	}

	var m models.Comment
	err = col.FindOne(ctx, filter).Decode(&m)
	if err != nil {
		repo.l.Errorf(ctx, "post.mongo.DetailComment.FindOne: %v", err)
		return models.Comment{}, err
	}

	return m, nil
}

func (repo impleRepository) GetOneComment(ctx context.Context, sc models.Scope, opts repository.GetOneCommentOptions) (models.Comment, error) {
	col := repo.getCommentCollection()

	filter, err := repo.buildGetOneCommentQuery(ctx, sc, opts.FilterComment)
	if err != nil {
		repo.l.Errorf(ctx, "post.mongo.GetOneComment.buildGetOneCommentQuery: %v", err)
		return models.Comment{}, err
	}

	var m models.Comment
	err = col.FindOne(ctx, filter).Decode(&m)
	if err != nil {
		repo.l.Errorf(ctx, "post.mongo.GetOneComment.FindOne: %v", err)
		return models.Comment{}, err
	}

	return m, nil
}

func (repo impleRepository) ListComment(ctx context.Context, sc models.Scope, opts repository.ListCommentOptions) ([]models.Comment, error) {
	col := repo.getCommentCollection()

	filter, err := repo.buildListCommentQuery(ctx, sc, opts)
	if err != nil {
		repo.l.Errorf(ctx, "post.mongo.ListComment.buildListCommentQuery: %v", err)
		return []models.Comment{}, err
	}

	cur, err := col.Find(ctx, filter)
	if err != nil {
		repo.l.Errorf(ctx, "post.mongo.ListComment.buildListCommentQuery: %v", err)
		return []models.Comment{}, err
	}

	var ms []models.Comment
	err = cur.All(ctx, ms)
	if err != nil {
		repo.l.Errorf(ctx, "post.mongo.ListComment.All: %v", err)
		return []models.Comment{}, err
	}

	return ms, nil
}

func (repo impleRepository) GetComment(ctx context.Context, sc models.Scope, opts repository.GetCommentOptions) ([]models.Comment, paginator.Paginator, error) {
	col := repo.getCommentCollection()

	filter, err := repo.buildGetCommentQuery(ctx, sc, opts)
	if err != nil {
		repo.l.Errorf(ctx, "post.mongo.GetComment.buildGetCommentQuery: %v", err)
		return []models.Comment{}, paginator.Paginator{}, err
	}

	cur, err := col.Find(ctx, filter, options.Find().
		SetLimit(opts.PagQuery.Limit).
		SetSkip(opts.PagQuery.Offset()))
	if err != nil {
		repo.l.Errorf(ctx, "post.mongo.GetComment.Find: %v", err)
		return []models.Comment{}, paginator.Paginator{}, err
	}

	var ms []models.Comment
	err = cur.All(ctx, &ms)
	if err != nil {
		repo.l.Errorf(ctx, "post.mongo.GetComment.All: %v", err)
		return []models.Comment{}, paginator.Paginator{}, err
	}

	total, err := col.CountDocuments(ctx, filter)
	if err != nil {
		repo.l.Errorf(ctx, "post.mongo.GetComment.CountDocuments: %v", err)
		return []models.Comment{}, paginator.Paginator{}, err
	}

	return ms, paginator.Paginator{
		Total:       total,
		Count:       int64(len(ms)),
		PerPage:     opts.PagQuery.Limit,
		CurrentPage: opts.PagQuery.Page,
	}, nil
}

func (repo impleRepository) UpdateComment(ctx context.Context, sc models.Scope, opts repository.UpdateCommentOptions) (models.Comment, error) {
	col := repo.getCommentCollection()

	m, filter, err := repo.buildUpdateCommentModels(ctx, sc, opts)
	if err != nil {
		repo.l.Errorf(ctx, "post.mongo.UpdateComment.buildUpdateCommentModels: %v", err)
		return models.Comment{}, nil
	}

	_, err = col.UpdateOne(ctx, filter, &m)
	if err != nil {
		repo.l.Errorf(ctx, "post.mongo.UpdateComment.UpdateOne: %v", err)
		return models.Comment{}, nil
	}

	return m, nil
}

func (repo impleRepository) DeleteComment(ctx context.Context, sc models.Scope, id string) error {
	col := repo.getCommentCollection()

	filter, err := repo.buildDetailCommentQuery(ctx, sc, id)
	if err != nil {
		repo.l.Errorf(ctx, "post.mongo.DeleteComment.buildDetailCommentQuery: %v", err)
		return err
	}

	_, err = col.DeleteOne(ctx, filter)
	if err != nil {
		repo.l.Errorf(ctx, "post.mongo.DeleteComment.DeleteOne: %v", err)
		return err
	}

	return nil
}
