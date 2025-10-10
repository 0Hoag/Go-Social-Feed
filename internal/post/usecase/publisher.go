package usecase

import (
	"context"

	rabb "github.com/hoag/go-social-feed/internal/feed/delivery/rabbitmq"
	"github.com/hoag/go-social-feed/internal/resource/notification"
)

func (uc impleUsecase) publishDeletePostRelationMsg(ctx context.Context, input rabb.DeletePostRelationMsg) error {
	msg := rabb.DeletePostRelationMsg{
		PostID: input.PostID,
	}

	err := uc.prod.PublishDeletePostRelationMsg(ctx, msg)
	if err != nil {
		uc.l.Errorf(ctx, "post.usecase.publishDeletePostRelationMsg: %v", err)
		return err
	}

	return nil
}

func (uc impleUsecase) publishPushNotiMsg(ctx context.Context, n notification.Notification) error {
	msg := rabb.PublishNotiMsg{
		Content:       n.Content,
		Heading:       n.Heading,
		UserIDs:       n.UserIDs,
		CreatedUserID: n.CreatedUserID,
		En: rabb.MultiLangObj{
			Heading: n.En.Heading,
			Content: n.En.Content,
		},
		Ja: rabb.MultiLangObj{
			Heading: n.Ja.Heading,
			Content: n.Ja.Content,
		},
		Data: rabb.NotiData{
			Data:     n.Data.Data,
			Activity: n.Data.Activity,
		},
		Source: n.Source,
	}

	err := uc.prod.PublishNotiMsg(ctx, msg)
	if err != nil {
		uc.l.Errorf(ctx, "event.usecase.publishPushNotiMsg: %v", err)
		return err
	}

	return nil
}
