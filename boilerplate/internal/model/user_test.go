package model_test

import (
	"testing"

	"boilerplate/internal/model"

	"github.com/stretchr/testify/require"
)

func TestUserIsValid(t *testing.T) {
	user := model.User{
		Name:  "John",
		Email: "john@example.com",
	}

	require.True(t, user.IsValid())
}

func TestUserIsInvalid(t *testing.T) {
	user := model.User{}

	require.False(t, user.IsValid())
}
