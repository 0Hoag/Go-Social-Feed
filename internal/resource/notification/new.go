package notification

import (
	"context"

	i18nPkg "github.com/hoag/go-social-feed/pkg/i18n"
	"github.com/hoag/go-social-feed/pkg/locale"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var GetNotiHeading = func(ctx context.Context, input GetNotiHeadingInput) (string, error) {
	var l string
	var ok bool
	if input.Lang != "" {
		l = input.Lang
	} else {
		l, ok = locale.GetLocaleFromContext(ctx)
		if !ok {
			return "", locale.ErrLocaleNotFound
		}
	}

	localizer := i18nPkg.NewLocalizer(l)

	switch input.From {
	case SourceNewsFeed:
		return localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "noti.newsfeed_header"}), nil
	}
	return "", nil
}

var GetNotiContent = func(ctx context.Context, input GetNotiContentInput) (string, error) {
	if _, ok := locale.SupportedLanguages[input.Lang]; !ok {
		return "", locale.ErrLocaleNotFound
	}

	localizer := i18nPkg.NewLocalizer(input.Lang)

	switch input.From {
	case SourcePostCreate:
		return localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "noti.post_created", TemplateData: map[string]interface{}{
			"TaggerName": input.TaggerName,
			"Content":    input.Content,
		}}), nil
	case SourcePostReactPost:
		return localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "noti.react_post", TemplateData: map[string]interface{}{
			"TaggerName": input.TaggerName,
			"Content":    input.Content,
		}}), nil
	case SourceTagUser:
		return localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "noti.tag_user", TemplateData: map[string]interface{}{
			"TaggerName": input.TaggerName,
			"Content":    input.Content,
		}}), nil
	}

	return "", nil
}
