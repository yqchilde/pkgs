package server

import (
	"github.com/yqchilde/gint/api/router"
	"github.com/yqchilde/gint/pkg/config"
	"github.com/yqchilde/gint/pkg/transport/http"
)

func NewHttpServer(c *config.Config) *http.Server {
	r := router.InitRouters()

	var opts []http.ServerOption
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	srv := http.NewServer(opts...)
	srv.Handler = r

	return srv
}
