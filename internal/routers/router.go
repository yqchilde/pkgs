package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/yqchilde/gin-skeleton/docs"
	"github.com/yqchilde/gin-skeleton/internal/handler/v1/application"
	"github.com/yqchilde/gin-skeleton/pkg/middleware"
)

func NewRouter() *gin.Engine {
	//r := gin.New()
	r := gin.Default()
	//r.Use(middleware.NoCache)
	//r.Use(middleware.Options)
	//r.Use(middleware.Logging())

	// swagger api docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Application center
	appApi := r.Group("/app")
	{
		appV1 := appApi.Group("/v1")
		{
			appV1.POST("/register", application.Register)
			appV1.POST("/login", application.Login)
			appV1.Use(middleware.JWT())
			{
				appV1.POST("/application", application.CreateApp)
				appV1.DELETE("/application/:id", application.DeleteApp)
			}
		}
	}

	return r
}
