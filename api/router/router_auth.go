package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yqchilde/gint/api/handler"
)

func InitAuthRouter(Router *gin.RouterGroup) {
	authRouter := Router.Group("auth")
	{
		authRouter.POST("casbin", handler.AddCasbinRule)
		authRouter.POST("jwt", handler.AddJwtBlackList)
	}
}
