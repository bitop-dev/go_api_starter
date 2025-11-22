package routes

import (
	"github.com/gofiber/fiber/v2"

	"go_api_starter/internal/http/handlers"
)

// Register mounts all API routes with versioning.
func Register(app *fiber.App, userHandler *handlers.UserHandler) {
	v1 := app.Group("/v1")

	v1.Get("/health", handlers.Health)

	users := v1.Group("/users")
	users.Get("/", userHandler.List)
	users.Post("/", userHandler.Create)
	users.Get(":id", userHandler.Get)
	users.Put(":id", userHandler.Update)
	users.Delete(":id", userHandler.Delete)
}
