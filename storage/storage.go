package storage

import (
	"context"
	"fasttrackquiz/types"
)

// Storage interface defines the behaviour of any Storage (memory, mysql, mongo..)
type Storage interface {
	Get(ctx context.Context) ([]types.QuizQuestion, error)
	Submit(ctx context.Context, req *types.QuizSubmitRequest) (*types.QuizResult, error)
}
