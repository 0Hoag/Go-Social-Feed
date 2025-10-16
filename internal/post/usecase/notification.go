package usecase

import (
	"context"

	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/internal/post"
	"github.com/hoag/go-social-feed/internal/resource/notification"
	"github.com/hoag/go-social-feed/pkg/locale"
	"github.com/hoag/go-social-feed/pkg/util"
)

// ============================================================================
// MAIN NOTIFICATION HANDLERS
// ============================================================================
func (uc impleUsecase) handleCreatePostNotification(ctx context.Context, sc models.Scope, p models.Post) error {
	user, err := uc.userUC.GetSessionUser(ctx, sc)
	if err != nil {
		uc.l.Errorf(ctx, "post.usecase.notification.handleCreatePostNotification.GetSessionUser: %v", err)
		return err
	}

	notiInput := getPostNotiContent{
		P:            p,
		Type:         notification.SourcePostCreate,
		Content:      p.Content,
		TaggerName:   user.Username,
		TaggedTarget: util.ObjectIDsToHex(p.TaggedTarget),
	}

	return uc.createAndPublishNotification(ctx, sc, notiInput)
}

// ============================================================================
// NOTIFICATION CREATION
// ============================================================================
func (uc impleUsecase) getPostNoti(ctx context.Context, sc models.Scope, input getPostNotiContent) (notification.Notification, error) {
	processedContent := util.TruncateHTMLString(input.Content, 15)

	langs := locale.SupportedLanguages
	headings := make(map[string]string)
	contents := make(map[string]string)

	for lang, ok := range langs {
		if !ok {
			continue
		}

		heading, err := notification.GetNotiHeading(ctx, notification.GetNotiHeadingInput{
			Lang: lang,
			From: notification.SourceNewsFeed,
		})
		if err != nil {
			uc.l.Errorf(ctx, "post.usecase.getPostNoti.GetNotiHeading: %v", err)
			return notification.Notification{}, err
		}

		headings[lang] = heading

		content, err := notification.GetNotiContent(ctx, notification.GetNotiContentInput{
			Lang:       lang,
			From:       input.Type,
			TaggerName: input.TaggerName,
			Content:    processedContent,
		})
		if err != nil {
			uc.l.Errorf(ctx, "post.usecase.getPostNoti.GetNotiContent: %v", err)
			return notification.Notification{}, err
		}

		headings[lang] = content

	}

	n := notification.Notification{
		Content:       contents[locale.ViLanguage],
		Heading:       headings[locale.ViLanguage],
		UserIDs:       input.TaggedTarget,
		CreatedUserID: sc.UserID,
		Data: notification.NotiData{
			Data: post.PublishNotiPostInput{
				PostID:     input.P.ID.Hex(),
				ReceiverID: input.P.AuthorID.Hex(),
				Type:       input.Type,
			},
			Activity: notification.ActivityPostDetail,
		},
		Source: notification.SourceNewsFeed,
	}

	if _, ok := langs[locale.EnLanguage]; ok {
		n.En = notification.MultiLangObj{
			Heading: headings[locale.EnLanguage],
			Content: contents[locale.EnLanguage],
		}
	}

	if _, ok := langs[locale.JaLanguage]; ok {
		n.Ja = notification.MultiLangObj{
			Heading: headings[locale.JaLanguage],
			Content: contents[locale.JaLanguage],
		}
	}

	return n, nil
}
