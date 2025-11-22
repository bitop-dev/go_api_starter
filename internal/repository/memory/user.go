package memory

import (
	"context"
	"sync"
	"time"

	domain "go_api_starter/internal/domain/user"
	pkgerrors "go_api_starter/pkg/errors"
)

// UserRepository is an in-memory repository for rapid prototyping.
type UserRepository struct {
	mu    sync.RWMutex
	seq   int64
	users map[int64]domain.User
}

// NewUserRepository constructs an empty in-memory repository.
func NewUserRepository() *UserRepository {
	return &UserRepository{users: make(map[int64]domain.User)}
}

// Create inserts a new user.
func (r *UserRepository) Create(ctx context.Context, user domain.User) (domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.seq++
	user.ID = r.seq
	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now()
	}
	r.users[user.ID] = user
	return user, nil
}

// Get fetches a user by ID.
func (r *UserRepository) Get(ctx context.Context, id int64) (domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	user, ok := r.users[id]
	if !ok {
		return domain.User{}, pkgerrors.New(pkgerrors.CodeNotFound, "user not found", nil)
	}
	return user, nil
}

// List returns all users.
func (r *UserRepository) List(ctx context.Context) ([]domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	users := make([]domain.User, 0, len(r.users))
	for _, u := range r.users {
		users = append(users, u)
	}
	return users, nil
}

// Update modifies an existing record.
func (r *UserRepository) Update(ctx context.Context, user domain.User) (domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, ok := r.users[user.ID]
	if !ok {
		return domain.User{}, pkgerrors.New(pkgerrors.CodeNotFound, "user not found", nil)
	}
	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now()
	}
	r.users[user.ID] = user
	return user, nil
}

// Delete removes a record.
func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.users[id]; !ok {
		return pkgerrors.New(pkgerrors.CodeNotFound, "user not found", nil)
	}
	delete(r.users, id)
	return nil
}
