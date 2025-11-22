package user

import (
	"context"
	"testing"
)

func TestCreateUserValidation(t *testing.T) {
	repo := newMockRepo()
	svc := NewService(repo)

	_, err := svc.CreateUser(context.Background(), "", "")
	if err == nil {
		t.Fatalf("expected validation error")
	}
}

// Minimal mock repository for validation tests.
type mockRepo struct{}

func newMockRepo() *mockRepo { return &mockRepo{} }

func (m *mockRepo) Create(ctx context.Context, u User) (User, error) { return u, nil }
func (m *mockRepo) Get(ctx context.Context, id int64) (User, error)  { return User{}, nil }
func (m *mockRepo) List(ctx context.Context) ([]User, error)         { return nil, nil }
func (m *mockRepo) Update(ctx context.Context, u User) (User, error) { return u, nil }
func (m *mockRepo) Delete(ctx context.Context, id int64) error       { return nil }
