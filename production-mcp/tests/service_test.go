package tests

import (
    "context"
    "testing"

    "production-mcp/internal/application/service"
    "production-mcp/internal/domain"
)

type mockRepo struct{}

func (m *mockRepo) GetService(ctx context.Context, name string) (*domain.Service, error) {
    return &domain.Service{Name: name}, nil
}

func (m *mockRepo) GetDependencies(ctx context.Context, name string) ([]string, error) {
    return []string{"payment-service"}, nil
}

func (m *mockRepo) GetDependents(ctx context.Context, name string) ([]string, error) {
    return []string{"booking-service"}, nil
}

func TestGetDependencies(t *testing.T) {
    svc := service.New(&mockRepo{})

    got, err := svc.GetDependencies(context.Background(), "booking-service")
    if err != nil {
        t.Fatal(err)
    }

    if len(got) != 1 || got[0] != "payment-service" {
        t.Fatalf("unexpected result: %#v", got)
    }
}
