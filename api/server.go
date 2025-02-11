package api

import (
	"context"
	"fasttrackquiz/storage"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	listenAddr string
	store      storage.Storage
	logger     *slog.Logger
	httpServer *http.Server
}

func NewServer(logger *slog.Logger, listenAddr string, store storage.Storage) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
		logger:     logger,
	}
}

func (s *Server) routes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/quiz", func(r chi.Router) {
			r.Get("/questions", s.handleFetchQuizQuestions)
			r.Post("/submit", s.handleSubmitAnswer)
		})
	})
	return r
}

func (s *Server) Start() error {
	s.logger.Info("Starting server", "listenAddr", s.listenAddr)
	s.httpServer = &http.Server{
		Addr:         s.listenAddr,
		Handler:      s.routes(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		s.logger.Error("Server encountered an error", "error", err)
		return err
	}

	s.logger.Info("Server closed normally")
	return nil
}

func (s *Server) Close(ctx context.Context) error {
	s.logger.Info("Shutting down server", "listenAddr", s.listenAddr)
	if s.httpServer == nil {
		s.logger.Info("httpServer is nil, nothing to shutdown")
		return nil
	}
	if err := s.httpServer.Shutdown(ctx); err != nil {
		s.logger.Error("Server forced to shutdown", "error", err)
		return err
	}
	s.logger.Info("Server shutdown complete")
	return nil
}
