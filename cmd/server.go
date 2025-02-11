package cmd

import (
	"context"
	"fasttrackquiz/api"
	"fasttrackquiz/storage"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

var server = &cobra.Command{
	Use:   "start-server",
	Short: "Start the quiz API server",
	Run: func(cmd *cobra.Command, args []string) {
		startServer()
	},
}

func startServer() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "3000"
	}

	logger := slog.Default()
	store := storage.NewMemoryStorage(logger)
	server := api.NewServer(logger, port, store)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.Start(); err != nil {
			logger.Error("Server failed", "error", err)
			cancel()
		}
	}()

	sig := <-sigChan
	logger.Info("Shutting down server", "signal", sig)

	closeCtx, closeCancel := context.WithTimeout(ctx, 5*time.Second)
	defer closeCancel()

	if err := server.Close(closeCtx); err != nil {
		logger.Error("Error during server shutdown", "error", err)
	} else {
		logger.Info("Server shutdown gracefully")
	}
}
