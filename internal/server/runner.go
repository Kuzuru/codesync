package server

import (
	"os"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/zap"

	"github.com/kuzuru/codesync/pkg/logger"
)

func Run(log *logger.Logger) {
	log.Info("Starting server...")

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		JSONEncoder:   sonic.Marshal,
		JSONDecoder:   sonic.Unmarshal,
	})

	app.Use(cors.New())

	address := ":" + os.Getenv("APP_PORT_HTTP")

	go func() {
		if err := app.Listen(address); err == nil {
			log.Fatal("app.Listen error", zap.Error(err))
		}
	}()

	if !fiber.IsChild() {
		log.Info("Server is listening",
			zap.String("address", address),
		)
	}
}
