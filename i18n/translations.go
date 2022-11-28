package i18n

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var uni = ut.New(en.New(), zh.New())

func Translations() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		language := ctx.GetHeader("language")
		trans, _ := uni.GetTranslator(language)
		v, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			switch language {
			case "zh":
				_ = zh_translations.RegisterDefaultTranslations(v, trans)
			case "en":
				_ = en_translations.RegisterDefaultTranslations(v, trans)
			default:
				_ = en_translations.RegisterDefaultTranslations(v, trans)
			}
			ctx.Set("trans", trans)
		}

		ctx.Next()
	}
}
