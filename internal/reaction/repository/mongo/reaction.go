package mongo

import (
	"context"
	"fmt"

	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/internal/reaction/repository"
	"github.com/hoag/go-social-feed/pkg/mongo"
	"github.com/hoag/go-social-feed/pkg/paginator"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	ReactionCollection = "reaction"
)

func (repo impleRepository) getReactionCollection() mongo.Collection {
	collName := fmt.Sprintf("%s", ReactionCollection)
	return repo.db.Collection(collName)
}

func (repo impleRepository) Create(ctx context.Context, sc models.Scope, opts repository.CreateOptions) (models.Reaction, error) {
	col := repo.getReactionCollection()

	m, err := repo.buildModels(ctx, sc, opts)
	if err != nil {
		repo.l.Errorf(ctx, "Reactions.mongo.Create.buildModels: %v", err)
		return models.Reaction{}, err
	}

	_, err = col.InsertOne(ctx, m)
	if err != nil {
		repo.l.Errorf(ctx, "Reactions.mogno.Create.InsertOne: %v", err)
		return models.Reaction{}, err
	}

	return m, nil
}

func (repo impleRepository) Detail(ctx context.Context, sc models.Scope, id string) (models.Reaction, error) {
	col := repo.getReactionCollection()

	filter, err := repo.buildDetailQuery(ctx, sc, id)
	if err != nil {
		repo.l.Errorf(ctx, "Reactions.mongo.Detail.buildDetailQuery: %v", err)
		return models.Reaction{}, err
	}

	var m models.Reaction
	err = col.FindOne(ctx, filter).Decode(&m)
	if err != nil {
		repo.l.Errorf(ctx, "Reactions.mongo.Detail.FindOne: %v", err)
		return models.Reaction{}, err
	}

	return m, nil
}

func (repo impleRepository) List(ctx context.Context, sc models.Scope, opts repository.ListOptions) ([]models.Reaction, error) {
	col := repo.getReactionCollection()

	filter, err := repo.buildListQuery(ctx, sc, opts)
	if err != nil {
		repo.l.Errorf(ctx, "Reactions.mongo.List.buildListQuery: %v", err)
		return []models.Reaction{}, err
	}

	cur, err := col.Find(ctx, filter)
	if err != nil {
		repo.l.Errorf(ctx, "Reactions.mongo.List.buildListQuery: %v", err)
		return []models.Reaction{}, err
	}

	var ms []models.Reaction
	err = cur.All(ctx, ms)
	if err != nil {
		repo.l.Errorf(ctx, "Reactions.mongo.List.All: %v", err)
		return []models.Reaction{}, err
	}

	return ms, nil
}

func (repo impleRepository) Get(ctx context.Context, sc models.Scope, opts repository.GetOptions) ([]models.Reaction, paginator.Paginator, error) {
	col := repo.getReactionCollection()

	filter, err := repo.buildGetQuery(ctx, sc, opts)
	if err != nil {
		repo.l.Errorf(ctx, "Reactions.mongo.Get.buildGetQuery: %v", err)
		return []models.Reaction{}, paginator.Paginator{}, err
	}

	cur, err := col.Find(ctx, filter, options.Find().
		SetLimit(opts.PagQuery.Limit).
		SetSkip(opts.PagQuery.Offset()))
	if err != nil {
		repo.l.Errorf(ctx, "Reactions.mongo.Get.Find: %v", err)
		return []models.Reaction{}, paginator.Paginator{}, err
	}

	var ms []models.Reaction
	err = cur.All(ctx, &ms)
	if err != nil {
		repo.l.Errorf(ctx, "Reactions.mongo.Get.All: %v", err)
		return []models.Reaction{}, paginator.Paginator{}, err
	}

	total, err := col.CountDocuments(ctx, filter)
	if err != nil {
		repo.l.Errorf(ctx, "Reactions.mongo.Get.CountDocuments: %v", err)
		return []models.Reaction{}, paginator.Paginator{}, err
	}

	return ms, paginator.Paginator{
		Total:       total,
		Count:       int64(len(ms)),
		PerPage:     opts.PagQuery.Limit,
		CurrentPage: opts.PagQuery.Page,
	}, nil
}

func (repo impleRepository) Delete(ctx context.Context, sc models.Scope, id string) error {
	col := repo.getReactionCollection()

	filter, err := repo.buildDetailQuery(ctx, sc, id)
	if err != nil {
		repo.l.Errorf(ctx, "Reactions.mongo.Delete.buildDetailQuery: %v", err)
		return err
	}

	_, err = col.DeleteOne(ctx, filter)
	if err != nil {
		repo.l.Errorf(ctx, "Reactions.mongo.Delete.DeleteOne: %v", err)
		return err
	}

	return nil
}
