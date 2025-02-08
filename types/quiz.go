package types

type Question struct {
	ID      int      `json:"id"`
	Text    string   `json:"text"`
	Options []Option `json:"options"`
}

type Option struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type Result struct {
	CorrectCount int     `json:"correct_count"`
	Percentile   float64 `json:"percentile"`
}

type SubmitRequest struct {
	Answers map[string]int `json:"answers"`
}
