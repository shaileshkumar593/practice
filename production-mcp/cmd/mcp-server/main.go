package main

import (
    "context"
    "log/slog"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/neo4j/neo4j-go-driver/v5/neo4j"
    "github.com/modelcontextprotocol/go-sdk/mcp"

    appservice "production-mcp/internal/application/service"
    "production-mcp/internal/config"
    graphrepo "production-mcp/internal/infrastructure/graph"
    mcpserver "production-mcp/internal/mcp"
)

func main() {
    logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
    slog.SetDefault(logger)

    cfg := config.Load()

    ctx, cancel := signal.NotifyContext(
        context.Background(),
        os.Interrupt,
        syscall.SIGTERM,
    )
    defer cancel()

    driver, err := neo4j.NewDriverWithContext(
        cfg.Neo4jURI,
        neo4j.BasicAuth(cfg.Neo4jUser, cfg.Neo4jPass, ""),
    )
    if err != nil {
        logger.Error("failed to create neo4j driver", "error", err)
        os.Exit(1)
    }

    defer func() {
        closeCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()
        driver.Close(closeCtx)
    }()

    pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()

    if err := driver.VerifyConnectivity(pingCtx); err != nil {
        logger.Error("neo4j connectivity failed", "error", err)
        os.Exit(1)
    }

    repository := graphrepo.New(driver)
    service := appservice.New(repository)
    server := mcpserver.New(service)

    logger.Info("starting MCP server", "environment", cfg.Environment)

    if err := server.Run(ctx, &mcp.StdioTransport{}); err != nil {
        logger.Error("MCP server stopped", "error", err)
        os.Exit(1)
    }
}
