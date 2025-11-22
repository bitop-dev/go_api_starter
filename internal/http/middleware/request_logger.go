package middleware

import (
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v2"
)

// RequestLogger logs basic request metadata using slog.
func RequestLogger(log *slog.Logger) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		start := time.Now()
		err := ctx.Next()
		duration := time.Since(start)

		log.InfoContext(ctx.Context(), "request",
			"method", ctx.Method(),
			"path", ctx.Path(),
			"status", ctx.Response().StatusCode(),
			"duration_ms", duration.Milliseconds(),
		)

		return err
	}
}
