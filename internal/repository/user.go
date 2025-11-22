package repository

import (
	"context"

	domain "go_api_starter/internal/domain/user"
)

// UserRepository exposes persistent operations for the user domain.
type UserRepository interface {
	Create(ctx context.Context, user domain.User) (domain.User, error)
	Get(ctx context.Context, id int64) (domain.User, error)
	List(ctx context.Context) ([]domain.User, error)
	Update(ctx context.Context, user domain.User) (domain.User, error)
	Delete(ctx context.Context, id int64) error
}
