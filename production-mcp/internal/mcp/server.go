package mcpserver

import (
    "github.com/modelcontextprotocol/go-sdk/mcp"

    appservice "production-mcp/internal/application/service"
)

func New(service *appservice.Service) *mcp.Server {
    server := mcp.NewServer(
        &mcp.Implementation{
            Name:    "production-knowledge-server",
            Version: "1.0.0",
        },
        nil,
    )

    tools := NewServiceTool(service)

    mcp.AddTool(server, &mcp.Tool{
        Name:        "get_service",
        Description: "Get information about a service.",
    }, tools.GetService)

    mcp.AddTool(server, &mcp.Tool{
        Name:        "get_dependencies",
        Description: "Find services that a service depends on.",
    }, tools.GetDependencies)

    mcp.AddTool(server, &mcp.Tool{
        Name:        "get_dependents",
        Description: "Find services that depend on a service.",
    }, tools.GetDependents)

    return server
}
