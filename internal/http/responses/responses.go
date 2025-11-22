package responses

import (
	"github.com/gofiber/fiber/v2"

	pkgerrors "go_api_starter/pkg/errors"
)

// JSON sends a standardized success response.
func JSON(ctx *fiber.Ctx, status int, payload interface{}) error {
	return ctx.Status(status).JSON(payload)
}

// Error sends a structured error response using the project schema.
func Error(ctx *fiber.Ctx, status int, err pkgerrors.APIError) error {
	return ctx.Status(status).JSON(fiber.Map{
		"error": fiber.Map{
			"code":    err.Code,
			"message": err.Message,
			"details": err.Details,
		},
	})
}
