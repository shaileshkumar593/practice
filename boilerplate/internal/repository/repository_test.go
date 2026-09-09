package repository_test

import (
	"context"
	"database/sql"
	"regexp"
	"testing"
	"time"

	"boilerplate/internal/model"
	"boilerplate/internal/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := repository.NewRepository(db)

	expectedUser := model.User{
		ID:        1,
		Name:      "John",
		Email:     "john@example.com",
		CreatedAt: time.Now(),
	}

	rows := sqlmock.NewRows(
		[]string{"id", "name", "email", "created_at"},
	).AddRow(
		expectedUser.ID,
		expectedUser.Name,
		expectedUser.Email,
		expectedUser.CreatedAt,
	)

	mock.ExpectQuery(
		regexp.QuoteMeta(`
			SELECT id, name, email, created_at
			FROM users
			WHERE id = $1
		`),
	).
		WithArgs(int64(1)).
		WillReturnRows(rows)

	got, err := repo.GetByID(
		context.Background(),
		1,
	)

	require.NoError(t, err)
	require.Equal(t, expectedUser.ID, got.ID)
	require.Equal(t, expectedUser.Name, got.Name)
	require.Equal(t, expectedUser.Email, got.Email)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetByIDNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := repository.NewRepository(db)

	mock.ExpectQuery(
		regexp.QuoteMeta(`
			SELECT id, name, email, created_at
			FROM users
			WHERE id = $1
		`),
	).
		WithArgs(int64(999)).
		WillReturnError(sql.ErrNoRows)

	_, err = repo.GetByID(
		context.Background(),
		999,
	)

	require.Error(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := repository.NewRepository(db)

	user := model.User{
		Name:  "John",
		Email: "john@example.com",
	}

	createdAt := time.Now()

	rows := sqlmock.NewRows(
		[]string{"id", "created_at"},
	).AddRow(
		int64(1),
		createdAt,
	)

	mock.ExpectQuery(
		regexp.QuoteMeta(`
			INSERT INTO users (name, email)
			VALUES ($1, $2)
			RETURNING id, created_at
		`),
	).
		WithArgs(user.Name, user.Email).
		WillReturnRows(rows)

	got, err := repo.Create(
		context.Background(),
		user,
	)

	require.NoError(t, err)
	require.Equal(t, int64(1), got.ID)
	require.Equal(t, user.Name, got.Name)
	require.Equal(t, user.Email, got.Email)

	require.NoError(t, mock.ExpectationsWereMet())
}
