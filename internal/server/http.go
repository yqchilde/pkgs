package server

import (
	"github.com/yqchilde/gin-skeleton/internal/routers"
	"github.com/yqchilde/gin-skeleton/pkg/conf"
	"github.com/yqchilde/gin-skeleton/pkg/transport/http"
)

func NewHttpServer(c *conf.Config) *http.Server {
	router := routers.NewRouter()

	var opts []http.ServerOption
	if c.HTTP.Network != "" {
		opts = append(opts, http.Network(c.HTTP.Network))
	}
	if c.HTTP.Addr != "" {
		opts = append(opts, http.Address(c.HTTP.Addr))
	}
	if c.HTTP.ReadTimeout != 0 {
		opts = append(opts, http.Timeout(c.HTTP.ReadTimeout))
	}
	if c.HTTP.WriteTimeout != 0 {
		opts = append(opts, http.Timeout(c.HTTP.WriteTimeout))
	}
	srv := http.NewServer(opts...)

	srv.Handler = router

	return srv
}
