package repository

import (
    "context"
    "production-mcp/internal/domain"
)

type GraphRepository interface {
    GetService(ctx context.Context, name string) (*domain.Service, error)
    GetDependencies(ctx context.Context, name string) ([]string, error)
    GetDependents(ctx context.Context, name string) ([]string, error)
}
