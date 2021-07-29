package router

import "github.com/gin-gonic/gin"

func InitRouters() *gin.Engine {
	var r = gin.Default()

	privateGroup := r.Group("v1")
	{
		InitAuthRouter(privateGroup)
	}

	return r
}
