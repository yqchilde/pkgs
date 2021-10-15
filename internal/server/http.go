package server

import (
	"github.com/yqchilde/gin-skeleton/internal/routers"
	"github.com/yqchilde/gin-skeleton/pkg/conf"
	"github.com/yqchilde/gin-skeleton/pkg/transport/http"
)

func NewHttpServer(c *conf.Config) *http.Server {
	router := routers.NewRouter()

	srv := http.NewServer(
		http.WithAddress(c.HTTP.Addr),
		http.WithReadTimeout(c.HTTP.ReadTimeout),
		http.WithWriteTimeout(c.HTTP.WriteTimeout),
	)

	srv.Handler = router

	return srv
}
