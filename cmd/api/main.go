package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"go_api_starter/internal/config"
	domain "go_api_starter/internal/domain/user"
	"go_api_starter/internal/http/handlers"
	"go_api_starter/internal/http/middleware"
	"go_api_starter/internal/http/routes"
	"go_api_starter/internal/repository/memory"
	"go_api_starter/pkg/logger"
)

func main() {
	cfg, err := config.Load(".")
	if err != nil {
		panic(fmt.Errorf("config: %w", err))
	}

	log := logger.New(cfg.Logging.Level, cfg.Logging.JSON)

	repo := memory.NewUserRepository()
	service := domain.NewService(repo)
	handler := handlers.NewUserHandler(service)

	app := fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})
	registerMiddleware(app, log)
	routes.Register(app, handler)
	registerDocs(app)

	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Info("starting server", "addr", addr)
	if err := app.Listen(addr); err != nil {
		log.Error("server stopped", "err", err)
		os.Exit(1)
	}
}

func registerMiddleware(app *fiber.App, log *slog.Logger) {
	app.Use(requestid.New())
	app.Use(middleware.ErrorHandler(log))
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(limiter.New(limiter.Config{Max: 100, Expiration: 60 * time.Second}))
	app.Use(middleware.RequestLogger(log))
}

func registerDocs(app *fiber.App) {
	app.Get("/docs", func(ctx *fiber.Ctx) error {
		return ctx.SendFile("internal/docs/openapi.yaml")
	})
}
