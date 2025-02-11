package storage

import (
	"context"
	"fasttrackquiz/types"
)

// Storage interface defines the behaviour of any Storage (memory, mysql, mongo..)
type Storage interface {
	GetAllQuestions(ctx context.Context) ([]types.QuizQuestion, error)
	SubmitAnswers(ctx context.Context, req *types.QuizSubmitRequest) (*types.QuizResult, error)
}
