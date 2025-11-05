package mongo

import (
	"context"

	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/internal/post/repository"
	"github.com/hoag/go-social-feed/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo impleRepository) buildCommentModels(ctx context.Context, sc models.Scope, opts repository.CreateCommentOptions) (models.Comment, error) {
	now := repo.clock()

	postID, err := primitive.ObjectIDFromHex(opts.PostID)
	if err != nil {
		repo.l.Errorf(ctx, "reaction.repository.buildCommentModels.ObjectIDFromHex: %v", err)
		return models.Comment{}, err
	}

	var attachments []models.Attachment
	if len(opts.Attach) > 0 {
		attachments = opts.Attach
	}

	comment := models.Comment{
		ID:          repo.db.NewObjectID(),
		PostID:      postID,
		AuthorID:    mongo.ObjectIDFromHexOrNil(sc.UserID),
		Content:     opts.Content,
		Attachments: attachments,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	return comment, nil
}

func (repo impleRepository) buildUpdateCommentModels(ctx context.Context, sc models.Scope, opts repository.UpdateCommentOptions) (models.Comment, bson.M, error) {
	now := repo.clock()

	set := bson.M{}

	if opts.Content != "" {
		set["content"] = opts.Content
	}

	if len(opts.Attach) > 0 {
		set["attachments"] = opts.Attach
	}

	set["updated_at"] = now

	return opts.Comment, set, nil
}
