package api

import (
	"bytes"
	"encoding/json"
	"fasttrackquiz/storage"
	"fasttrackquiz/types"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleFetchQuizQuestions(t *testing.T) {
	store := storage.NewMemoryStorage(nil)
	server := NewServer(slog.Default(), ":3000", store)

	req := httptest.NewRequest("GET", "/quiz-question", nil)
	w := httptest.NewRecorder()

	server.handleFetchQuizQuestions(w, req)
	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response struct {
		Data []types.QuizQuestion `json:"data"`
	}

	err := json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)
	assert.NotEmpty(t, response.Data)
	assert.Len(t, response.Data, 5)
}

func TestHandleSubmitAnswer(t *testing.T) {
	store := storage.NewMemoryStorage(nil)
	server := NewServer(slog.Default(), ":3000", store)

	payload := &types.QuizSubmitRequest{
		QuizAnswers: map[int]int{
			1: 1,
			2: 2,
			3: 3,
			4: 1,
			5: 2,
		},
	}

	jsonPayload, _ := json.Marshal(payload)
	req := httptest.NewRequest("POST", "/submit", bytes.NewReader(jsonPayload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	server.handleSubmitAnswer(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response struct {
		Data types.QuizResult `json:"data"`
	}

	err := json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, 2, response.Data.CorrectCount)
	assert.GreaterOrEqual(t, response.Data.Percentile, 0.0)
	assert.LessOrEqual(t, response.Data.Percentile, 100.0)
}
