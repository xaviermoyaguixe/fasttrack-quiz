package api

import (
	"context"
	"fasttrackquiz/storage"
	"log/slog"
	"net/http"
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
	mux := http.NewServeMux()
	mux.HandleFunc("/questions", s.handleFetchQuizQuestions)
	mux.HandleFunc("/submit", s.handleSubmitAnswer)
	return mux
}

func (s *Server) Start() error {
	s.logger.Info("starting server...")
	s.httpServer = &http.Server{
		Addr:    s.listenAddr,
		Handler: s.routes(),
	}

	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (s *Server) Close(ctx context.Context) error {
	s.logger.Info("Shutting down server...")
	if err := s.httpServer.Shutdown(ctx); err != nil {
		s.logger.Error("Server forced to shutdown", "error", err)
		return err
	}
	s.logger.Info("Server shutdown complete")
	return nil
}
