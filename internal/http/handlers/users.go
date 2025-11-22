package handlers

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"

	domain "go_api_starter/internal/domain/user"
	"go_api_starter/internal/http/responses"
	pkgerrors "go_api_starter/pkg/errors"
)

// FiberUserService exposes domain methods in a Fiber-friendly way.
type FiberUserService interface {
	CreateUser(ctx context.Context, email, name string) (domain.User, error)
	GetUser(ctx context.Context, id int64) (domain.User, error)
	ListUsers(ctx context.Context) ([]domain.User, error)
	UpdateUser(ctx context.Context, id int64, email, name string) (domain.User, error)
	DeleteUser(ctx context.Context, id int64) error
}

// UserHandler manages user endpoints.
type UserHandler struct {
	service FiberUserService
}

// NewUserHandler constructs a handler with the provided service.
func NewUserHandler(service FiberUserService) *UserHandler {
	return &UserHandler{service: service}
}

// Create handles POST /v1/users.
type createUserRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (h *UserHandler) Create(ctx *fiber.Ctx) error {
	var req createUserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return pkgerrors.New(pkgerrors.CodeInvalidPayload, "invalid request body", err.Error())
	}

	user, err := h.service.CreateUser(ctx.Context(), req.Email, req.Name)
	if err != nil {
		return err
	}

	return responses.JSON(ctx, fiber.StatusCreated, user)
}

// Get handles GET /v1/users/:id.
func (h *UserHandler) Get(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		return pkgerrors.New(pkgerrors.CodeInvalidPayload, "invalid user id", nil)
	}

	user, err := h.service.GetUser(ctx.Context(), id)
	if err != nil {
		return err
	}

	return responses.JSON(ctx, fiber.StatusOK, user)
}

// List handles GET /v1/users.
func (h *UserHandler) List(ctx *fiber.Ctx) error {
	users, err := h.service.ListUsers(ctx.Context())
	if err != nil {
		return err
	}

	return responses.JSON(ctx, fiber.StatusOK, users)
}

// Update handles PUT /v1/users/:id.
type updateUserRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (h *UserHandler) Update(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		return pkgerrors.New(pkgerrors.CodeInvalidPayload, "invalid user id", nil)
	}
	var req updateUserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return pkgerrors.New(pkgerrors.CodeInvalidPayload, "invalid request body", err.Error())
	}

	user, err := h.service.UpdateUser(ctx.Context(), id, req.Email, req.Name)
	if err != nil {
		return err
	}

	return responses.JSON(ctx, fiber.StatusOK, user)
}

// Delete handles DELETE /v1/users/:id.
func (h *UserHandler) Delete(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		return pkgerrors.New(pkgerrors.CodeInvalidPayload, "invalid user id", nil)
	}

	if err := h.service.DeleteUser(ctx.Context(), id); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
