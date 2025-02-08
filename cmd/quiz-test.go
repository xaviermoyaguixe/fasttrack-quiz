package cmd

import (
	"fasttrackquiz/api"
	"fasttrackquiz/storage"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

const listenAddr = ":3000"

var quizTestCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the quiz and answer the questions",
	RunE: func(cmd *cobra.Command, args []string) error {
		return startQuiz()
	},
}

func startQuiz() error {
	store := storage.NewMemoryStorage()
	server := api.NewServer(listenAddr, store)
	fmt.Printf("server running on port: %s", listenAddr)
	log.Fatal(server.Start())
	return nil
}
