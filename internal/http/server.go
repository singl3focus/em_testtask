package http

import (
	"context"
	"net/http"
	"time"
)

const (
	ReadTimeout       = time.Second * 20
	ReadHeaderTimeout = time.Second * 10
	WriteTimeout      = time.Second * 20
)

type Server struct {
	server *http.Server
}

func NewServer(port string, handler http.Handler) *Server {
	return &Server{
		server: &http.Server{
			Addr:              ":" + port,
			Handler:           handler,
			ReadTimeout:       ReadTimeout,
			ReadHeaderTimeout: ReadHeaderTimeout,
			WriteTimeout:      WriteTimeout,
		},
	}
}

func (s *Server) Start() {
	if err := s.server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
