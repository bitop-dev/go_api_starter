package handlers

import (
	"github.com/gofiber/fiber/v2"

	"go_api_starter/internal/http/responses"
)

// Health returns service readiness.
func Health(ctx *fiber.Ctx) error {
	return responses.JSON(ctx, fiber.StatusOK, fiber.Map{"status": "ok"})
}
