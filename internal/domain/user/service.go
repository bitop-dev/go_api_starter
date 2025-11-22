package user

import (
	"context"
	"fmt"
	"time"

	pkgerrors "go_api_starter/pkg/errors"
)

// User represents a persisted user record.
type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

// Repository defines the storage interface for users.
type Repository interface {
	Create(ctx context.Context, user User) (User, error)
	Get(ctx context.Context, id int64) (User, error)
	List(ctx context.Context) ([]User, error)
	Update(ctx context.Context, user User) (User, error)
	Delete(ctx context.Context, id int64) error
}

// Service orchestrates user operations.
type Service struct {
	repo Repository
}

// NewService constructs a Service with the provided repository.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// CreateUser adds a new user after basic validation.
func (s *Service) CreateUser(ctx context.Context, email, name string) (User, error) {
	if email == "" || name == "" {
		return User{}, pkgerrors.New(pkgerrors.CodeInvalidPayload, "email and name are required", nil)
	}

	newUser := User{Email: email, Name: name, CreatedAt: time.Now()}
	created, err := s.repo.Create(ctx, newUser)
	if err != nil {
		return User{}, fmt.Errorf("create user: %w", err)
	}
	return created, nil
}

// GetUser fetches a user by ID.
func (s *Service) GetUser(ctx context.Context, id int64) (User, error) {
	u, err := s.repo.Get(ctx, id)
	if err != nil {
		return User{}, err
	}
	return u, nil
}

// ListUsers returns all users.
func (s *Service) ListUsers(ctx context.Context) ([]User, error) {
	return s.repo.List(ctx)
}

// UpdateUser updates an existing record.
func (s *Service) UpdateUser(ctx context.Context, id int64, email, name string) (User, error) {
	if email == "" || name == "" {
		return User{}, pkgerrors.New(pkgerrors.CodeInvalidPayload, "email and name are required", nil)
	}

	u := User{ID: id, Email: email, Name: name}
	updated, err := s.repo.Update(ctx, u)
	if err != nil {
		return User{}, fmt.Errorf("update user: %w", err)
	}
	return updated, nil
}

// DeleteUser deletes a record.
func (s *Service) DeleteUser(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
