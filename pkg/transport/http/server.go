package http

import (
	"context"
	"errors"
	"net"
	"net/http"
	"time"

	"github.com/yqchilde/gint/pkg/logger"
)

type Server struct {
	*http.Server
	lis     net.Listener
	network string
	address string
	timeout time.Duration
	log     logger.Logger
}

func NewServer(opts ...ServerOption) *Server {
	srv := &Server{
		network: "tcp",
		address: ":8080",
		timeout: time.Second,
		log:     logger.GetLogger(),
	}

	for _, o := range opts {
		o(srv)
	}

	srv.Server = &http.Server{Handler: srv}
	return srv
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.ServeHTTP(w, r)
}

func (s *Server) Start() error {
	lis, err := net.Listen(s.network, s.address)
	if err != nil {
		return err
	}
	s.lis = lis
	s.log.Info("[HTTP] Server is listening on: %s", lis.Addr().String())
	if err := s.Serve(lis); !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (s *Server) Stop() error {
	s.log.Info("[HTTP] Server is stopping")
	return s.Shutdown(context.Background())
}
