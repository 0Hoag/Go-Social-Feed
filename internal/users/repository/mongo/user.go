package mongo

import (
	"context"
	"fmt"

	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/internal/users/repository"
	"github.com/hoag/go-social-feed/pkg/mongo"
	"github.com/hoag/go-social-feed/pkg/paginator"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	userCollection = "users"
)

func (repo impleRepository) getUserCollection() mongo.Collection {
	collName := fmt.Sprintf("%s", userCollection)
	return repo.db.Collection(collName)
}

func (repo impleRepository) Create(ctx context.Context, sc models.Scope, opts repository.CreateOptions) (models.User, error) {
	col := repo.getUserCollection()

	m, err := repo.buildModels(ctx, sc, opts)
	if err != nil {
		repo.l.Errorf(ctx, "users.mongo.Create.buildModels: %v", err)
		return models.User{}, err
	}

	_, err = col.InsertOne(ctx, m)
	if err != nil {
		repo.l.Errorf(ctx, "users.mogno.Create.InsertOne: %v", err)
		return models.User{}, err
	}

	return m, nil
}

func (repo impleRepository) Detail(ctx context.Context, sc models.Scope, id string) (models.User, error) {
	col := repo.getUserCollection()

	filter, err := repo.buildDetailQuery(ctx, sc, id)
	if err != nil {
		repo.l.Errorf(ctx, "users.mongo.Detail.buildDetailQuery: %v", err)
		return models.User{}, err
	}

	var m models.User
	err = col.FindOne(ctx, filter).Decode(&m)
	if err != nil {
		repo.l.Errorf(ctx, "users.mongo.Detail.FindOne: %v", err)
		return models.User{}, err
	}

	return m, nil
}

func (repo impleRepository) List(ctx context.Context, sc models.Scope, opts repository.ListOptions) ([]models.User, error) {
	col := repo.getUserCollection()

	filter, err := repo.buildListQuery(ctx, sc, opts)
	if err != nil {
		repo.l.Errorf(ctx, "users.mongo.List.buildListQuery: %v", err)
		return []models.User{}, err
	}

	cur, err := col.Find(ctx, filter)
	if err != nil {
		repo.l.Errorf(ctx, "users.mongo.List.buildListQuery: %v", err)
		return []models.User{}, err
	}

	var ms []models.User
	err = cur.All(ctx, ms)
	if err != nil {
		repo.l.Errorf(ctx, "users.mongo.List.All: %v", err)
		return []models.User{}, err
	}

	return ms, nil
}

func (repo impleRepository) Get(ctx context.Context, sc models.Scope, opts repository.GetOptions) ([]models.User, paginator.Paginator, error) {
	col := repo.getUserCollection()

	filter, err := repo.buildGetQuery(ctx, sc, opts)
	if err != nil {
		repo.l.Errorf(ctx, "users.mongo.Get.buildGetQuery: %v", err)
		return []models.User{}, paginator.Paginator{}, err
	}

	cur, err := col.Find(ctx, filter, options.Find().
		SetLimit(opts.PagQuery.Limit).
		SetSkip(opts.PagQuery.Offset()))
	if err != nil {
		repo.l.Errorf(ctx, "users.mongo.Get.Find: %v", err)
		return []models.User{}, paginator.Paginator{}, err
	}

	var ms []models.User
	err = cur.All(ctx, ms)
	if err != nil {
		repo.l.Errorf(ctx, "users.mongo.Get.All: %v", err)
		return []models.User{}, paginator.Paginator{}, err
	}

	total, err := col.CountDocuments(ctx, filter)
	if err != nil {
		repo.l.Errorf(ctx, "users.mongo.Get.CountDocuments: %v", err)
		return []models.User{}, paginator.Paginator{}, err
	}

	return ms, paginator.Paginator{
		Total:       total,
		Count:       int64(len(ms)),
		PerPage:     opts.PagQuery.Limit,
		CurrentPage: opts.PagQuery.Page,
	}, nil
}

func (repo impleRepository) Update(ctx context.Context, sc models.Scope, opts repository.UpdateOptions) error {
	col := repo.getUserCollection()

	filter, err := repo.buildDetailQuery(ctx, sc, opts.User.ID.Hex())
	if err != nil {
		repo.l.Errorf(ctx, "users.mongo.Update.buildUpdateModels: %v", err)
		return err
	}

	update, err := repo.buildUpdateModels(ctx, sc, opts)
	if err != nil {
		repo.l.Errorf(ctx, "users.mongo.Update.buildUpdateModels: %v", err)
		return err
	}

	_, err = col.UpdateOne(ctx, filter, update)
	if err != nil {
		repo.l.Errorf(ctx, "users.mongo.Update.UpdateOne: %v", err)
		return err
	}

	return nil
}

func (repo impleRepository) Delete(ctx context.Context, sc models.Scope, id string) error {
	col := repo.getUserCollection()

	filter, err := repo.buildDetailQuery(ctx, sc, id)
	if err != nil {
		repo.l.Errorf(ctx, "users.mongo.Delete.buildDetailQuery: %v", err)
		return err
	}

	_, err = col.DeleteOne(ctx, filter)
	if err != nil {
		repo.l.Errorf(ctx, "users.mongo.Delete.DeleteOne: %v", err)
		return err
	}

	return nil
}
