package graph

import (
    "context"
    "fmt"

    "github.com/neo4j/neo4j-go-driver/v5/neo4j"

    "production-mcp/internal/domain"
)

type Neo4jRepository struct {
    driver neo4j.DriverWithContext
}

func New(driver neo4j.DriverWithContext) *Neo4jRepository {
    return &Neo4jRepository{driver: driver}
}

func (r *Neo4jRepository) GetService(ctx context.Context, name string) (*domain.Service, error) {
    session := r.driver.NewSession(ctx, neo4j.SessionConfig{
        AccessMode: neo4j.AccessModeRead,
    })
    defer session.Close(ctx)

    result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
        result, err := tx.Run(ctx, `
            MATCH (s:Service {name: $name})
            RETURN s.id AS id, s.name AS name, s.version AS version,
                   s.status AS status, s.owner AS owner
        `, map[string]any{"name": name})
        if err != nil {
            return nil, err
        }

        if !result.Next(ctx) {
            return nil, fmt.Errorf("service %q not found", name)
        }

        return result.Record().AsMap(), nil
    })
    if err != nil {
        return nil, err
    }

    data := result.(map[string]any)
    return &domain.Service{
        ID:      stringValue(data["id"]),
        Name:    stringValue(data["name"]),
        Version: stringValue(data["version"]),
        Status:  stringValue(data["status"]),
        Owner:   stringValue(data["owner"]),
    }, nil
}

func (r *Neo4jRepository) GetDependencies(ctx context.Context, name string) ([]string, error) {
    return r.relatedServices(ctx, `
        MATCH (s:Service {name: $name})-[:DEPENDS_ON*1..3]->(d:Service)
        RETURN DISTINCT d.name AS name
    `, name)
}

func (r *Neo4jRepository) GetDependents(ctx context.Context, name string) ([]string, error) {
    return r.relatedServices(ctx, `
        MATCH (d:Service)-[:DEPENDS_ON*1..3]->(s:Service {name: $name})
        RETURN DISTINCT d.name AS name
    `, name)
}

func (r *Neo4jRepository) relatedServices(ctx context.Context, query, name string) ([]string, error) {
    session := r.driver.NewSession(ctx, neo4j.SessionConfig{
        AccessMode: neo4j.AccessModeRead,
    })
    defer session.Close(ctx)

    result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
        result, err := tx.Run(ctx, query, map[string]any{"name": name})
        if err != nil {
            return nil, err
        }

        var services []string
        for result.Next(ctx) {
            if value, ok := result.Record().Get("name"); ok {
                if name, ok := value.(string); ok {
                    services = append(services, name)
                }
            }
        }
        return services, result.Err()
    })
    if err != nil {
        return nil, err
    }
    return result.([]string), nil
}

func stringValue(v any) string {
    if s, ok := v.(string); ok {
        return s
    }
    return ""
}
