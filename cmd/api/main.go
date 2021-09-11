package main

import (
	"github.com/gin-gonic/gin"

	"github.com/yqchilde/gint/internal/dao"
	"github.com/yqchilde/gint/internal/model"
	"github.com/yqchilde/gint/internal/server"
	"github.com/yqchilde/gint/internal/service"
	"github.com/yqchilde/gint/pkg/app"
	"github.com/yqchilde/gint/pkg/config"
	"github.com/yqchilde/gint/pkg/logger"
	"github.com/yqchilde/gint/pkg/redis"
)

func main() {
	// init config
	cfg, err := config.Init()
	if err != nil {
		panic(err)
	}

	// init logger, mysql and redis
	logger.Init(&cfg.Logger)
	model.Init(&cfg.MySQL)
	redis.Init(&cfg.Redis)

	// init service
	service.New(cfg, dao.New(model.GetDB(), cfg))

	// set gin mode
	gin.SetMode(cfg.Server.Mode)

	// init app
	a := app.New(cfg,
		app.WithName("gint"),
		app.Server(
			server.NewHttpServer(cfg),
		),
	)

	if err := a.Run(); err != nil {
		panic(err)
	}
}
