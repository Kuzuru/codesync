package main

import (
	"github.com/kuzuru/codesync/pkg/logger"
)

func main() {
	log, err := logger.Init()
	if err != nil {
		panic(err)
	}

	log.Info("Zap Logger initialized successfully")
}
