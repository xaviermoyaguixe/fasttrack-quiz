package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var quizTestCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the quiz and answer the questions",
	RunE: func(cmd *cobra.Command, args []string) error {
		return startQuiz()
	},
}

func startQuiz() error {
	fmt.Println("COMMAND IS WORKING")
	return nil
}
