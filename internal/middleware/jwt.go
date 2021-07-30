package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/yqchilde/gint/internal/service"
	"github.com/yqchilde/gint/pkg/app"
	"github.com/yqchilde/gint/pkg/errcode"
)

func JwtBlackList() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtToken := c.GetHeader("Authorization")
		ok, err := service.AuthSvc.CheckJwtIsBlackList(c, jwtToken)
		if err == nil && ok {
			app.NewResponse().Error(c, errcode.ErrInvalidToken)
			c.Abort()
			return
		}
	}
}
