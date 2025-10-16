package usecase

import (
	"context"

	"github.com/hoag/go-social-feed/internal/models"
)

// Helper function to create and publish notification
func (uc impleUsecase) createAndPublishNotification(ctx context.Context, sc models.Scope, input getPostNotiContent) error {
	noti, err := uc.getPostNoti(ctx, sc, input)
	if err != nil {
		uc.l.Errorf(ctx, "post.usecase.util.getPostNoti: %v", err)
		return err
	}

	err = uc.publishPushNotiMsg(ctx, noti)
	if err != nil {
		uc.l.Errorf(ctx, "post.usecase.util.publishPushNotiMsg: %v", err)
		return err
	}

	return nil
}
