package app

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type ValidError struct {
	Key string
	Err string
}

func (v *ValidError) Error() string {
	return strings.ReplaceAll(v.Err, v.Key, v.Key[strings.Index(v.Key, ".")+1:])
}

type ValidErrors []*ValidError

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

func init() {
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validate.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}

// BindAndValid bind form and verify params
func BindAndValid(ctx *gin.Context, v interface{}) (bool, ValidErrors) {
	var errs ValidErrors
	if err := ctx.ShouldBind(v); err != nil {
		v := ctx.Value("trans")
		trans, _ := v.(ut.Translator)
		valErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			return false, errs
		}

		for field, err := range valErrors.Translate(trans) {
			errs = append(errs, &ValidError{
				Key: field,
				Err: err,
			})
		}

		return false, errs
	}

	return true, nil
}
