package api

import (
	"context"
	"encoding/json"
	"fasttrackquiz/types"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
)

func (s *Server) handleFetchQuizQuestions(w http.ResponseWriter, r *http.Request) {
	reqID := middleware.GetReqID(r.Context())

	s.logger.Info("Fetching quiz questions", "method", r.Method, "request_id", reqID)

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	questions, err := s.store.Get(ctx)
	if err != nil {
		s.logger.Error("Error retrieving questions from store", "request_id", reqID, "err", err.Error())
		http.Error(w, "Failed to retrieve quiz questions", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, questions, "succesfully fetched quiz questions")
}

func (s *Server) handleSubmitAnswer(w http.ResponseWriter, r *http.Request) {
	reqID := middleware.GetReqID(r.Context())

	s.logger.Info("Starting quiz submission", "method", r.Method, "request_id", reqID)

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	defer r.Body.Close()

	var req types.QuizSubmitRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.logger.Error("Invalid JSON payload", "request_id", reqID, "err", err.Error())
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	if len(req.QuizAnswers) == 0 {
		s.logger.Error("Empty quiz submission received", "request_id", reqID)
		http.Error(w, "Quiz submission cannot be empty", http.StatusBadRequest)
		return
	}

	quizResult, err := s.store.Submit(ctx, &req)
	if err != nil {
		s.logger.Error("Error processing quiz submission", "request_id", reqID, "err", err.Error())
		http.Error(w, "Failed to process quiz submission", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, quizResult, "successfully submitted quiz answers")
}

func writeJSON(w http.ResponseWriter, status int, data any, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]any{
		"data":    data,
		"message": message,
	})
}
