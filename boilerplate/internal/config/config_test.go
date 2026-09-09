package config_test

import (
	"os"
	"testing"

	"boilerplate/internal/config"

	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	t.Setenv("APP_NAME", "test-app")
	t.Setenv("HTTP_PORT", "8081")
	t.Setenv(
		"DATABASE_URL",
		"postgres://postgres:postgres@localhost:5432/test",
	)

	cfg := config.Load()

	require.Equal(t, "test-app", cfg.AppName)
	require.Equal(t, "8081", cfg.HTTPPort)
	require.Equal(
		t,
		"postgres://postgres:postgres@localhost:5432/test",
		"postgres://postgres:postgres@localhost:5432/test",
	)
}

func TestLoadDefaults(t *testing.T) {
	os.Unsetenv("APP_NAME")
	os.Unsetenv("HTTP_PORT")

	cfg := config.Load()

	require.NotEmpty(t, cfg.AppName)
	require.NotEmpty(t, cfg.HTTPPort)
}
