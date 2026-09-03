package service

import (
    "context"
    "fmt"
    "strings"

    "production-mcp/internal/domain"
    "production-mcp/internal/repository"
)

type Service struct {
    repo repository.GraphRepository
}

func New(repo repository.GraphRepository) *Service {
    return &Service{repo: repo}
}

func (s *Service) GetService(ctx context.Context, name string) (*domain.Service, error) {
    name = strings.TrimSpace(name)
    if name == "" {
        return nil, fmt.Errorf("service name cannot be empty")
    }
    return s.repo.GetService(ctx, name)
}

func (s *Service) GetDependencies(ctx context.Context, name string) ([]string, error) {
    name = strings.TrimSpace(name)
    if name == "" {
        return nil, fmt.Errorf("service name cannot be empty")
    }
    return s.repo.GetDependencies(ctx, name)
}

func (s *Service) GetDependents(ctx context.Context, name string) ([]string, error) {
    name = strings.TrimSpace(name)
    if name == "" {
        return nil, fmt.Errorf("service name cannot be empty")
    }
    return s.repo.GetDependents(ctx, name)
}
