package main

import (
	"github.com/gin-gonic/gin"

	"github.com/yqchilde/gin-skeleton/internal/server"
	"github.com/yqchilde/gin-skeleton/internal/store"
	ginSkeleton "github.com/yqchilde/gin-skeleton/pkg/app"
	"github.com/yqchilde/gin-skeleton/pkg/conf"
	logger "github.com/yqchilde/gin-skeleton/pkg/log"
	"github.com/yqchilde/gin-skeleton/pkg/redis"
)

func main() {
	// init config
	cfg, err := conf.Init()
	if err != nil {
		panic(err)
	}

	// init component
	logger.Init(&cfg.Logger)
	store.Init(&cfg.MySQL)
	redis.Init(&cfg.Redis)

	gin.SetMode(conf.Conf.App.Mode)

	app := ginSkeleton.New(
		cfg,
		ginSkeleton.WithName(cfg.App.Name),
		ginSkeleton.WithLogger(logger.GetLogger()),
		ginSkeleton.Server(
			server.NewHttpServer(conf.Conf),
		),
	)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
