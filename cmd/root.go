package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "quiz-cli",
	Short: "A CLI tool for running the iGaming quiz",
	Long:  "quiz-cli is a command-line interface that allows users to start and take a quiz.",
}

func Execute() error {
	rootCmd.AddCommand(server)
	rootCmd.AddCommand(client)
	return rootCmd.Execute()
}
