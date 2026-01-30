package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/pandaxgh/ecom-backend/internal/config"
	"github.com/pandaxgh/ecom-backend/internal/database"
)

type Server struct {
	Config     *config.Config
	DB         *database.DB
	httpServer *http.Server
}

func New(cfg *config.Config) (*Server, error) {
	db := database.Connect(cfg.DB.URL)

	server := &Server{
		Config: cfg,
		DB:     db,
	}
	return server, nil
}

func (s *Server) SetupHttpServer(handler http.Handler) {
	s.httpServer = &http.Server{
		Addr:         ":" + s.Config.HTTP.Port,
		Handler:      handler,
		ReadTimeout:  time.Duration(s.Config.HTTP.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(s.Config.HTTP.WriteTimeout) * time.Second,
	}
}

func (s *Server) Start() error {
	if s.httpServer == nil {
		return errors.New("Http Server Not Initialized")
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) ShutDown(ctx context.Context) error {
	if err := s.httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("Failed to shutdown http server: %w", err)
	}

	if s.DB != nil && s.DB.Pool != nil {
		s.DB.Pool.Close()
	}

	return nil

}
