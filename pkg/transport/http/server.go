package http

import (
	"context"
	"errors"
	"net"
	"net/http"
	"time"

	"github.com/yqchilde/gin-skeleton/pkg/log"
)

type Server struct {
	*http.Server
	lis     net.Listener
	network string
	address string
	timeout time.Duration
	log     log.Logger
}

func NewServer(opts ...ServerOption) *Server {
	srv := &Server{
		network: "tcp",
		address: ":8080",
		timeout: time.Second,
		log:     log.GetLogger(),
	}
	for _, o := range opts {
		o(srv)
	}

	srv.Server = &http.Server{
		Handler: srv,
	}
	return srv
}

func (s *Server) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	s.ServeHTTP(resp, req)
}

func (s *Server) Start(ctx context.Context) error {
	lis, err := net.Listen(s.network, s.address)
	if err != nil {
		return err
	}
	s.lis = lis
	s.log.Infof("[HTTP] server is listening on: %s", lis.Addr().String())
	if err := s.Serve(lis); !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.log.Info("[HTTP] server is stopping")
	return s.Shutdown(ctx)
}
