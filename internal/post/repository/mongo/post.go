package mongo

import (
	"context"
	"fmt"

	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/internal/post/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	postCollection = "post"
)

func (repo impleRepository) getPostCollection() mongo.Collection {
	collName := fmt.Sprintf("%s", postCollection)
	return *repo.db.Collection(collName)
}

func (repo impleRepository) Create(ctx context.Context, sc models.Scope, opts repository.CreateOptions) (models.Post, error) {
	col := repo.getPostCollection()

	p := repo.buildModels(ctx, sc, opts)

}
