package endpoint_test

import (
	"context"
	"errors"
	"testing"

	"boilerplate/internal/endpoint"
	"boilerplate/internal/model"

	userreq "boilerplate/internal/request/user"
	userresp "boilerplate/internal/response/user"

	"github.com/stretchr/testify/require"
)

// --------------------------------------------------
// Mock Service
// --------------------------------------------------

type mockService struct {
	user model.User
	err  error
}

func (m *mockService) GetUser(
	ctx context.Context,
	id int64,
) (model.User, error) {
	if m.err != nil {
		return model.User{}, m.err
	}

	return m.user, nil
}

func (m *mockService) CreateUser(
	ctx context.Context,
	name string,
	email string,
) (model.User, error) {
	if m.err != nil {
		return model.User{}, m.err
	}

	return m.user, nil
}

// --------------------------------------------------
// MakeGetUserEndpoint - Success
// --------------------------------------------------

func TestMakeGetUserEndpoint_Success(t *testing.T) {
	expectedUser := model.User{
		ID:    1,
		Name:  "John",
		Email: "john@example.com",
	}

	mockSvc := &mockService{
		user: expectedUser,
	}

	ep := endpoint.MakeGetUserEndpoint(mockSvc)

	request := userreq.GetUserRequest{
		ID: 1,
	}

	response, err := ep(
		context.Background(),
		request,
	)

	require.NoError(t, err)
	require.NotNil(t, response)

	resp, ok := response.(userresp.GetUserResponse)

	require.True(t, ok)

	require.Equal(
		t,
		expectedUser,
		resp.User,
	)

	require.Empty(t, resp.Error)
}

// --------------------------------------------------
// MakeGetUserEndpoint - Error
// --------------------------------------------------

func TestMakeGetUserEndpoint_Error(t *testing.T) {
	expectedErr := errors.New("user not found")

	mockSvc := &mockService{
		err: expectedErr,
	}

	ep := endpoint.MakeGetUserEndpoint(mockSvc)

	request := userreq.GetUserRequest{
		ID: 999,
	}

	response, err := ep(
		context.Background(),
		request,
	)

	require.NoError(t, err)
	require.NotNil(t, response)

	resp, ok := response.(userresp.GetUserResponse)

	require.True(t, ok)

	require.Equal(
		t,
		expectedErr.Error(),
		resp.Error,
	)

	require.Equal(
		t,
		model.User{},
		resp.User,
	)
}

// --------------------------------------------------
// MakeCreateUserEndpoint - Success
// --------------------------------------------------

func TestMakeCreateUserEndpoint_Success(t *testing.T) {
	expectedUser := model.User{
		ID:    1,
		Name:  "John",
		Email: "john@example.com",
	}

	mockSvc := &mockService{
		user: expectedUser,
	}

	ep := endpoint.MakeCreateUserEndpoint(mockSvc)

	request := userreq.CreateUserRequest{
		Name:  "John",
		Email: "john@example.com",
	}

	response, err := ep(
		context.Background(),
		request,
	)

	require.NoError(t, err)
	require.NotNil(t, response)

	resp, ok := response.(userresp.CreateUserResponse)

	require.True(t, ok)

	require.Equal(
		t,
		expectedUser,
		resp.User,
	)

	require.Empty(t, resp.Error)
}

// --------------------------------------------------
// MakeCreateUserEndpoint - Error
// --------------------------------------------------

func TestMakeCreateUserEndpoint_Error(t *testing.T) {
	expectedErr := errors.New("database error")

	mockSvc := &mockService{
		err: expectedErr,
	}

	ep := endpoint.MakeCreateUserEndpoint(mockSvc)

	request := userreq.CreateUserRequest{
		Name:  "John",
		Email: "john@example.com",
	}

	response, err := ep(
		context.Background(),
		request,
	)

	require.NoError(t, err)
	require.NotNil(t, response)

	resp, ok := response.(userresp.CreateUserResponse)

	require.True(t, ok)

	require.Equal(
		t,
		expectedErr.Error(),
		resp.Error,
	)

	require.Equal(
		t,
		model.User{},
		resp.User,
	)
}
