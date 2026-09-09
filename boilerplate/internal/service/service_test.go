package service_test

import (
	"context"
	"errors"
	"testing"

	"boilerplate/internal/model"
	"boilerplate/internal/service"

	"github.com/stretchr/testify/require"
)

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

	return user, nil
}

func TestGetUser(t *testing.T) {
	expectedUser := model.User{
		ID:    1,
		Name:  "John",
		Email: "john@example.com",
	}

	repo := &mockRepository{
		user: expectedUser,
	}

	svc := service.NewService(repo)

	user, err := svc.GetUser(
		context.Background(),
		1,
	)

	require.NoError(t, err)
	require.Equal(t, expectedUser, user)
}

func TestGetUserRepositoryError(t *testing.T) {
	expectedErr := errors.New("database error")

	repo := &mockRepository{
		err: expectedErr,
	}

	svc := service.NewService(repo)

	user, err := svc.GetUser(
		context.Background(),
		1,
	)

	require.Error(t, err)
	require.ErrorIs(t, err, expectedErr)
	require.Equal(t, model.User{}, user)
}

func TestCreateUser(t *testing.T) {
	repo := &mockRepository{}

	svc := service.NewService(repo)

	user, err := svc.CreateUser(
		context.Background(),
		"John",
		"john@example.com",
	)

	require.NoError(t, err)

	require.Equal(t, "John", user.Name)
	require.Equal(t, "john@example.com", user.Email)
}

func TestCreateUserRepositoryError(t *testing.T) {
	expectedErr := errors.New("database error")

	repo := &mockRepository{
		err: expectedErr,
	}

	svc := service.NewService(repo)

	user, err := svc.CreateUser(
		context.Background(),
		"John",
		"john@example.com",
	)

	require.Error(t, err)
	require.ErrorIs(t, err, expectedErr)
	require.Equal(t, model.User{}, user)
}
