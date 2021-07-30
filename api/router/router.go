package router

import (
	"github.com/gin-gonic/gin"

	mw "github.com/yqchilde/gint/internal/middleware"
)

func InitRouters() *gin.Engine {
	var r = gin.Default()

	publicGroup := r.Group("v1")
	{
		InitPublicRouter(publicGroup)
	}

	privateGroup := r.Group("v1")
	privateGroup.Use(mw.JwtBlackList(), mw.Casbin())
	{
		InitAuthRouter(privateGroup)
		InitUserRouter(privateGroup)
	}

	return r
}
