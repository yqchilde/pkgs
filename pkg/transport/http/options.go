package http

import (
	"time"

	"github.com/yqchilde/gint/pkg/transport"
)

var _ transport.Server = (*Server)(nil)

type ServerOption func(*Server)

func Address(addr string) ServerOption {
	return func(s *Server) {
		s.address = addr
	}
}

func Timeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.timeout = timeout
	}
}
