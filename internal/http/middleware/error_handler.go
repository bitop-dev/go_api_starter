package middleware

import (
	"log/slog"
	nethttp "net/http"

	"github.com/gofiber/fiber/v2"

	"go_api_starter/internal/http/responses"
	pkgerrors "go_api_starter/pkg/errors"
)

// ErrorHandler converts errors into structured responses.
func ErrorHandler(log *slog.Logger) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		err := ctx.Next()
		if err == nil {
			return nil
		}

		apiErr, ok := err.(pkgerrors.APIError)
		if !ok {
			log.ErrorContext(ctx.Context(), "unexpected error", "err", err)
			apiErr = pkgerrors.New(pkgerrors.CodeInternal, "internal server error", nil)
			return responses.Error(ctx, nethttp.StatusInternalServerError, apiErr)
		}

		status := statusFromCode(apiErr.Code)
		log.WarnContext(ctx.Context(), "handled error", "code", apiErr.Code, "path", ctx.Path())
		return responses.Error(ctx, status, apiErr)
	}
}

func statusFromCode(code pkgerrors.Code) int {
	switch code {
	case pkgerrors.CodeInvalidPayload:
		return nethttp.StatusBadRequest
	case pkgerrors.CodeNotFound:
		return nethttp.StatusNotFound
	default:
		return nethttp.StatusInternalServerError
	}
}
