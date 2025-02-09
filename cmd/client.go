package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fasttrackquiz/types"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

const (
	// In a real world scenario this goes in ENV.
	serverAddr     = "http://localhost:3000"
	endpointFetch  = serverAddr + "/questions"
	endpointSubmit = serverAddr + "/submit"
	// In a real world scenario this goes in ENV.
)

var client = &cobra.Command{
	Use:   "start-client",
	Short: "Start the quiz and answer the questions",
	RunE: func(cmd *cobra.Command, args []string) error {
		return startClient()
	},
}

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func startClient() error {
	questions, err := fetchQuestions()
	if err != nil {
		fmt.Println("Error fetching questions:", err)
		return err
	}

	answers := collectUserAnswers(questions)

	result, err := submitAnswers(answers)
	if err != nil {
		fmt.Println("Error submitting answers:", err)
		return err
	}

	fmt.Printf("\n You answered %d out of %d questions correctly.\n", result.CorrectCount, len(questions))
	fmt.Printf(" Performance: You were better than %.0f%% of all quizzers.\n", result.Percentile*100)
	return nil
}

func makeRequest(method, url string, payload any, response any) error {
	var reqBody io.Reader
	if payload != nil {
		jsonData, err := json.Marshal(payload)
		if err != nil {
			return fmt.Errorf("failed to marshal payload: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("read body response failed: %w", err)
		}
		return fmt.Errorf("server error: %s (status %d)", string(body), resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return fmt.Errorf("failed to parse response: %w", err)
	}

	return nil
}

func fetchQuestions() ([]types.QuizQuestion, error) {
	var response struct {
		Success bool                 `json:"success"`
		Data    []types.QuizQuestion `json:"data"`
		Message string               `json:"message"`
	}

	if err := makeRequest("GET", endpointFetch, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func collectUserAnswers(questions []types.QuizQuestion) map[int]int {
	answers := make(map[int]int)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n Welcome to the iGaming Quiz! Answer the following questions:")

	for _, q := range questions {
		fmt.Printf("\n Question %d: %s\n", q.ID, q.Text)
		for idx, opt := range q.QuizOptions {
			fmt.Printf("  %d) %s\n", idx+1, opt.Text)
		}

		for {
			fmt.Printf("➡️  Select your answer (1-%d): ", len(q.QuizOptions))
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input. Try again.")
				continue
			}
			input = strings.TrimSpace(input)
			choice, err := strconv.Atoi(input)
			if err != nil || choice < 1 || choice > len(q.QuizOptions) {
				fmt.Println("Invalid selection, please try again.")
				continue
			}

			selectedOption := q.QuizOptions[choice-1]
			answers[q.ID] = selectedOption.ID
			break
		}
	}

	return answers
}

func submitAnswers(answers map[int]int) (*types.QuizResult, error) {
	submitPayload := types.QuizSubmitRequest{QuizAnswers: answers}

	var response struct {
		Success bool             `json:"success"`
		Data    types.QuizResult `json:"data"`
		Message string           `json:"message"`
	}

	if err := makeRequest("POST", endpointSubmit, submitPayload, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}
