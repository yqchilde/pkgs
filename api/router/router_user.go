package router

import "github.com/gin-gonic/gin"

func InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	{
		userRouter.POST("")
	}
}
