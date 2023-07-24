package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"redirect_api/internal/config"
	"time"
)

// server class
type Server struct {
	srv               *http.Server
	ServerErrorNotify chan error
}

// server object
func NewServer(cfg *config.Config, router *gin.Engine) *Server {
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

// run server while sending errors to error channel
func (s *Server) Run() {
	s.ServerErrorNotify <- s.srv.ListenAndServe()
}

// accept errors into channel
func (s *Server) ServerErrNotify() <-chan error {
	return s.ServerErrorNotify
}

// shutdown function
func (s *Server) Shutdown() error {
	return s.srv.Shutdown(context.Background())
}
