package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

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

	// init pprof server
	go func() {
		if err := http.ListenAndServe(conf.Conf.App.PprofPort, http.DefaultServeMux); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen ListenAndServe for Pprof, err: %s", err.Error())
		}
	}()

	app := ginSkeleton.New(
		cfg,
		ginSkeleton.WithName(cfg.App.Name),
		ginSkeleton.WithVersion(cfg.App.Version),
		ginSkeleton.WithLogger(logger.GetLogger()),
		ginSkeleton.Server(
			server.NewHttpServer(conf.Conf),
		),
	)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
