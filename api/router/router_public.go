package router

import (
	"github.com/gin-gonic/gin"

	"github.com/yqchilde/gint/api/handler"
)

func InitPublicRouter(Router *gin.RouterGroup) {
	publicRouter := Router.Group("public")
	{
		publicRouter.POST("signUp", handler.SignUp)
		publicRouter.POST("signIn", handler.SignIn)
	}
}
