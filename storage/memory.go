package storage

import (
	"fasttrackquiz/types"
	"sync"
)

type MemoryStorage struct {
	questions []types.Question
	mu        sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		questions: loadQuestions(),
	}
}

func (s *MemoryStorage) Get() ([]types.Question, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	questionsCopy := make([]types.Question, len(s.questions))
	copy(questionsCopy, s.questions)
	return questionsCopy, nil
}

func (s *MemoryStorage) Submit(*types.SubmitRequest) error {
	return nil
}

func loadQuestions() []types.Question {
	return []types.Question{
		{
			ID:   1,
			Text: "In online casino games, what does RTP stand for?",
			Options: []types.Option{
				{ID: 1, Text: "Return to Player"},
				{ID: 2, Text: "Real Time Payment"},
				{ID: 3, Text: "Return to Profit"},
			},
		},
		{
			ID:   2,
			Text: "Which type of slot game is most popular among players?",
			Options: []types.Option{
				{ID: 1, Text: "Video Slots"},
				{ID: 2, Text: "Classic Slots"},
				{ID: 3, Text: "Multi-Line Slots"},
			},
		},
		{
			ID:   3,
			Text: "What does RNG mean in the context of iGaming?",
			Options: []types.Option{
				{ID: 1, Text: "Random Number Generator"},
				{ID: 2, Text: "Real Network Gain"},
				{ID: 3, Text: "Remote Navigation Guide"},
			},
		},
	}
}
