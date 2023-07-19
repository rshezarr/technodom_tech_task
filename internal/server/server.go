package server

import (
	"context"
	"net/http"
	"redirect_api/internal/config"
	"time"
)

type Server struct {
	srv               *http.Server
	ServerErrorNotify chan error
}

func NewServer(cfg *config.Config, router *http.ServeMux) *Server {
	return &Server{
		srv: &http.Server{
			Addr:           cfg.API.Port,
			ReadTimeout:    time.Duration(cfg.API.Timeout) * time.Second,
			WriteTimeout:   time.Duration(cfg.API.Timeout) * time.Second,
			MaxHeaderBytes: cfg.API.MaxHeaderBytes << 20,
			Handler:        router,
		},
		ServerErrorNotify: make(chan error, 1),
	}
}

func (s *Server) Run() {
	s.ServerErrorNotify <- s.srv.ListenAndServe()
}

func (s *Server) ServerErrNotify() <-chan error {
	return s.ServerErrorNotify
}

func (s *Server) Shutdown() error {
	return s.srv.Shutdown(context.Background())
}
