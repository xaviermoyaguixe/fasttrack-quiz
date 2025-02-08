package api

import (
	"encoding/json"
	"fasttrackquiz/storage"
	"log"
	"net/http"
)

type Server struct {
	listenAddr string
	store      storage.Storage
}

func NewServer(listenAddr string, store storage.Storage) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/quiz-question", s.handleGetQuestion)
	return http.ListenAndServe(s.listenAddr, nil)
}

func (s *Server) handleGetQuestion(w http.ResponseWriter, r *http.Request) {
	question, _ := s.store.Get()
	log.Print("getting quiz question")
	json.NewEncoder(w).Encode(question)
}

func (s *Server) handleSubmitAnswer(w http.ResponseWriter, r *http.Request) {
	question := s.store.Submit(nil)
	log.Print("getting quiz question")
	json.NewEncoder(w).Encode(question)
}
