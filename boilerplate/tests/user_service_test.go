package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"boilerplate/internal/model"
	"boilerplate/internal/repository"
	"boilerplate/internal/service"

	"github.com/stretchr/testify/require"
)

// --------------------------------------------------
// Mock Repository
// --------------------------------------------------

type mockRepository struct {
	user model.User
	err  error
}

func (m *mockRepository) GetByID(
	ctx context.Context,
	id int64,
) (model.User, error) {

	if m.err != nil {
		return model.User{}, m.err
	}

	return m.user, nil
}

func (m *mockRepository) Create(
	ctx context.Context,
	user model.User,
) (model.User, error) {

	if m.err != nil {
		return model.User{}, m.err
	}

	user.ID = 1
	user.CreatedAt = time.Now()

	return user, nil
}

// --------------------------------------------------
// Test: Get User
// --------------------------------------------------

func TestGetUser(t *testing.T) {

	repo := &mockRepository{
		user: model.User{
			ID:    1,
			Name:  "Alice",
			Email: "alice@example.com",
		},
	}

	svc := service.NewService(repo)

	user, err := svc.GetUser(
		context.Background(),
		1,
	)

	require.NoError(t, err)

	require.Equal(
		t,
		int64(1),
		user.ID,
	)

	require.Equal(
		t,
		"Alice",
		user.Name,
	)

	require.Equal(
		t,
		"alice@example.com",
		user.Email,
	)
}

// --------------------------------------------------
// Test: Invalid User ID
// --------------------------------------------------

func TestGetUserInvalidID(t *testing.T) {

	repo := &mockRepository{}

	svc := service.NewService(repo)

	_, err := svc.GetUser(
		context.Background(),
		0,
	)

	require.Error(t, err)

	require.ErrorIs(
		t,
		err,
		service.ErrInvalidUser,
	)
}

// --------------------------------------------------
// Test: User Not Found
// --------------------------------------------------

func TestGetUserNotFound(t *testing.T) {

	repo := &mockRepository{
		err: repository.ErrUserNotFound,
	}

	svc := service.NewService(repo)

	_, err := svc.GetUser(
		context.Background(),
		1,
	)

	require.Error(t, err)

	require.ErrorIs(
		t,
		err,
		repository.ErrUserNotFound,
	)
}

// --------------------------------------------------
// Test: Create User Validation
// --------------------------------------------------

func TestCreateUserValidation(t *testing.T) {

	repo := &mockRepository{}

	svc := service.NewService(repo)

	_, err := svc.CreateUser(
		context.Background(),
		"",
		"bad-email",
	)

	require.Error(t, err)

	require.ErrorIs(
		t,
		err,
		service.ErrInvalidUser,
	)
}

// --------------------------------------------------
// Test: Create User Success
// --------------------------------------------------

func TestCreateUser(t *testing.T) {

	repo := &mockRepository{}

	svc := service.NewService(repo)

	user, err := svc.CreateUser(
		context.Background(),
		"Alice",
		"alice@example.com",
	)

	require.NoError(t, err)

	require.Equal(
		t,
		int64(1),
		user.ID,
	)

	require.Equal(
		t,
		"Alice",
		user.Name,
	)

	require.Equal(
		t,
		"alice@example.com",
		user.Email,
	)

	require.False(
		t,
		user.CreatedAt.IsZero(),
	)
}

// --------------------------------------------------
// Test: Create User Repository Error
// --------------------------------------------------

func TestCreateUserRepositoryError(t *testing.T) {

	expectedErr := errors.New("database error")

	repo := &mockRepository{
		err: expectedErr,
	}

	svc := service.NewService(repo)

	_, err := svc.CreateUser(
		context.Background(),
		"Alice",
		"alice@example.com",
	)

	require.Error(t, err)

	require.ErrorIs(
		t,
		err,
		expectedErr,
	)
}
