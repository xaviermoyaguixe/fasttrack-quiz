package storage

import (
	"context"
	"errors"
	"fasttrackquiz/types"
	"log/slog"
	"math/rand"
	"sync"
	"time"
)

type MemoryStorage struct {
	questions []types.QuizQuestion
	scores    []int
	logger    *slog.Logger

	mu sync.RWMutex //locks are not strictly necessary for this quiz-test, but are useful if scaling to multiple concurrent users.
}

func NewMemoryStorage(logger *slog.Logger) *MemoryStorage {
	return &MemoryStorage{
		questions: loadQuestions(),
		scores:    generateFakeScores(50),
		logger:    logger,
	}

}

func (ms *MemoryStorage) Get(ctx context.Context) ([]types.QuizQuestion, error) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	questionsCopy := append([]types.QuizQuestion{}, ms.questions...)
	return questionsCopy, nil
}

func (ms *MemoryStorage) Submit(ctx context.Context, req *types.QuizSubmitRequest) (*types.QuizResult, error) {
	if req == nil || len(req.QuizAnswers) == 0 {
		ms.logger.Error("Received empty or nil quiz submission")
		return nil, errors.New("no answers submitted")
	}

	ms.mu.Lock()
	defer ms.mu.Unlock()

	correctCount := 0
	for _, q := range ms.questions {
		select {
		case <-ctx.Done():
			ms.logger.Warn("Quiz submission canceled",
				"reason", ctx.Err())
			return nil, ctx.Err()
		default:
		}

		if userAnswer, exists := req.QuizAnswers[q.ID]; exists {
			if isCorrect := userAnswer == q.CorrectOptionID; isCorrect {
				correctCount++
			}
		}
	}

	ms.scores = append(ms.scores, correctCount)

	totalUsers := len(ms.scores)
	below := 0
	for _, score := range ms.scores {
		if score < correctCount {
			below++
		}
	}
	percentile := float64(below) / float64(totalUsers)

	return &types.QuizResult{
		CorrectCount: correctCount,
		Percentile:   percentile,
	}, nil
}

func loadQuestions() []types.QuizQuestion {
	return []types.QuizQuestion{
		{
			ID:   1,
			Text: "In online casino games, what does RTP stand for?",
			QuizOptions: []types.QuizOption{
				{ID: 1, Text: "Return to Player"},
				{ID: 2, Text: "Real Time Payment"},
				{ID: 3, Text: "Return to Profit"},
			},
			CorrectOptionID: 1,
		},
		{
			ID:   2,
			Text: "Which type of slot game is most popular among players?",
			QuizOptions: []types.QuizOption{
				{ID: 1, Text: "Classic Slots"},
				{ID: 2, Text: "Video Slots"},
				{ID: 3, Text: "Multi-Line Slots"},
			},
			CorrectOptionID: 2,
		},
		{
			ID:   3,
			Text: "What does RNG mean in the context of iGaming?",
			QuizOptions: []types.QuizOption{
				{ID: 1, Text: "Random Number Generator"},
				{ID: 2, Text: "Real Network Gain"},
				{ID: 3, Text: "Remote Navigation Guide"},
			},
			CorrectOptionID: 1,
		},
		{
			ID:   4,
			Text: "Which game offers the highest RTP (Return to Player) on average?",
			QuizOptions: []types.QuizOption{
				{ID: 1, Text: "Slots"},
				{ID: 2, Text: "Roulette"},
				{ID: 3, Text: "Blackjack"},
			},
			CorrectOptionID: 3,
		},
		{
			ID:   5,
			Text: "What is a progressive jackpot?",
			QuizOptions: []types.QuizOption{
				{ID: 1, Text: "A jackpot that increases over time"},
				{ID: 2, Text: "A fixed jackpot amount"},
				{ID: 3, Text: "A bonus that resets every day"},
			},
			CorrectOptionID: 1,
		},
	}
}

func generateFakeScores(numPlayers int) []int {
	rand.Seed(time.Now().UnixNano())
	var scores []int
	for i := 0; i < numPlayers; i++ {
		scores = append(scores, rand.Intn(6))
	}
	return scores
}
