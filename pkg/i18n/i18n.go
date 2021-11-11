package i18n

import (
	"log"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pelletier/go-toml"
	"golang.org/x/text/language"
)

type Config struct {
	FilePath []string `mapstructure:"file-path"`
}

type languageFunc func(language, translationID string, args ...interface{}) string

var (
	// T is translated function
	T languageFunc

	// DefaultLanguage is the default language
	DefaultLanguage = "en"
)

func Init(cfg *Config) {
	T = func(lang, translationId string, args ...interface{}) string {
		if lang == "" {
			lang = DefaultLanguage
		}
		if lang == "cn" {
			lang = "zh"
		}

		bundle := i18n.NewBundle(language.English)
		bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

		for _, filePath := range cfg.FilePath {
			if _, err := bundle.LoadMessageFile(filePath); err != nil {
				log.Panicf("the %s file was not found", filePath)
			}
		}

		loc := i18n.NewLocalizer(bundle, lang)
		trans, err := loc.Localize(&i18n.LocalizeConfig{
			MessageID: translationId,
		})
		if err != nil {
			return translationId
		}
		return trans
	}
}
