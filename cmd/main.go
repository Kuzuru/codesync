package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"go.uber.org/zap"

	"github.com/kuzuru/codesync/internal/server"
	"github.com/kuzuru/codesync/pkg/logger"
)

func main() {
	log := logger.New()

	defer log.Sync()

	log.Info("Zap Logger initialized successfully")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file",
			zap.String("message", err.Error()),
		)
	}

	log.Info(".env file loaded successfully")

	server.Run(log)

	// Waiting for exit signal from OS
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)

	<-quit
}
