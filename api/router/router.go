package router

import "github.com/gin-gonic/gin"

func InitRouters() *gin.Engine {
	var r = gin.Default()

	return r
}
