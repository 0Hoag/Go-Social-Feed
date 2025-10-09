package i18n

import (
	"embed"
	"encoding/json"
	"io/fs"
	"path/filepath"
	"strings"
	"sync"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var (
	bundle *i18n.Bundle
	once   sync.Once
)

//go:embed translations/*
var translationsFS embed.FS

func InitI18n() {
	once.Do(func() {
		bundle = i18n.NewBundle(language.English)
		bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

		// Load translation files
		loadTranslationFiles(translationsFS, "translations")
	})
}

func loadTranslationFiles(fsys embed.FS, dir string) {
	err := fs.WalkDir(fsys, dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(d.Name(), ".json") {
			data, err := fs.ReadFile(fsys, path)
			if err != nil {
				return err
			}
			bundle.ParseMessageFileBytes(data, filepath.Base(path))
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func NewLocalizer(lang string) *i18n.Localizer {
	return i18n.NewLocalizer(bundle, lang)
}
