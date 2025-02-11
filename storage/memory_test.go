package storage

import (
	"context"
	"fasttrackquiz/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetQuestions(t *testing.T) {
	store := NewMemoryStorage(nil)

	questions, err := store.GetAllQuestions(context.Background())
	assert.NoError(t, err)
	assert.NotEmpty(t, questions)
	assert.Len(t, questions, 5)
}

func TestSubmitAnswers(t *testing.T) {
	store := NewMemoryStorage(nil)

	req := &types.QuizSubmitRequest{
		QuizAnswers: map[int]int{
			1: 1,
			2: 2,
			3: 3,
			4: 1,
			5: 2,
		},
	}

	result, err := store.SubmitAnswers(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, 2, result.CorrectCount)
	assert.GreaterOrEqual(t, result.Percentile, 0.0)
	assert.LessOrEqual(t, result.Percentile, 100.0)
}
