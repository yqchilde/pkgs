package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/yqchilde/gin-skeleton/pkg/app"
	"github.com/yqchilde/gin-skeleton/pkg/errcode"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c, err := app.ParseRequest(ctx)
		if err != nil {
			app.NewResponse().Error(ctx, errcode.ErrInvalidToken)
			ctx.Abort()
			return
		}

		ctx.Set("user_id", c.UserID)
		ctx.Next()
	}
}
