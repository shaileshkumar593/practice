package mcpserver

import (
    "context"

    "github.com/modelcontextprotocol/go-sdk/mcp"

    appservice "production-mcp/internal/application/service"
)

type ServiceInput struct {
    Service string `json:"service" jsonschema:"service name"`
}

type ServiceTool struct {
    service *appservice.Service
}

func NewServiceTool(service *appservice.Service) *ServiceTool {
    return &ServiceTool{service: service}
}

func (t *ServiceTool) GetService(
    ctx context.Context,
    req *mcp.CallToolRequest,
    input ServiceInput,
) (*mcp.CallToolResult, any, error) {
    service, err := t.service.GetService(ctx, input.Service)
    if err != nil {
        return nil, nil, err
    }
    return nil, service, nil
}

func (t *ServiceTool) GetDependencies(
    ctx context.Context,
    req *mcp.CallToolRequest,
    input ServiceInput,
) (*mcp.CallToolResult, any, error) {
    dependencies, err := t.service.GetDependencies(ctx, input.Service)
    if err != nil {
        return nil, nil, err
    }
    return nil, map[string]any{
        "service":      input.Service,
        "dependencies": dependencies,
    }, nil
}

func (t *ServiceTool) GetDependents(
    ctx context.Context,
    req *mcp.CallToolRequest,
    input ServiceInput,
) (*mcp.CallToolResult, any, error) {
    dependents, err := t.service.GetDependents(ctx, input.Service)
    if err != nil {
        return nil, nil, err
    }
    return nil, map[string]any{
        "service":   input.Service,
        "dependents": dependents,
    }, nil
}
