package types

type QuizQuestion struct {
	ID              int          `json:"id"`
	Text            string       `json:"text"`
	QuizOptions     []QuizOption `json:"quiz_options"`
	CorrectOptionID int          `json:"correct_option_id"`
}

type QuizOption struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type QuizResult struct {
	CorrectCount int     `json:"correct_count"`
	Percentile   float64 `json:"percentile"`
}

type QuizSubmitRequest struct {
	QuizAnswers map[int]int `json:"quiz_answers"`
}
