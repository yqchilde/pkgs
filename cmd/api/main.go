package main

import (
	"github.com/gin-gonic/gin"

	"github.com/yqchilde/gin-skeleton/internal/server"
	"github.com/yqchilde/gin-skeleton/internal/store"
	"github.com/yqchilde/gin-skeleton/pkg/app"
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

	_app := app.New(
		cfg,
		app.WithName(cfg.App.Name),
		app.WithLogger(logger.GetLogger()),
		app.Server(
			server.NewHttpServer(conf.Conf),
		),
	)

	if err := _app.Run(); err != nil {
		panic(err)
	}
}
