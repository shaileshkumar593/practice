# Production MCP Server - Go + Knowledge Graph

Production-oriented reference architecture for an MCP server that exposes
service-knowledge tools backed by Neo4j.

## Architecture

AI Agent
  -> MCP Client
  -> MCP Server
  -> Application Service
  -> Graph Repository
  -> Neo4j / Knowledge Graph

## Tools

- get_service
- get_dependencies
- get_dependents

## Run locally

1. Install Go.
2. Set Neo4j environment variables:

   NEO4J_URI=bolt://localhost:7687
   NEO4J_USER=neo4j
   NEO4J_PASSWORD=change-me
   ENVIRONMENT=development

3. Download dependencies:

   go mod tidy

4. Run:

   go run ./cmd/mcp-server

## Important

The sample uses STDIO transport for easy local execution. For a remote
production deployment, use the Streamable HTTP transport provided by the
version of the official Go MCP SDK that you pin.

Authentication, authorization, rate limiting, audit logging, metrics,
OpenTelemetry, and secret management should be integrated with your
organization's production infrastructure before deployment.

## Example graph

(:Service {name:"booking-service"})
(:Service {name:"payment-service"})
(:Service {name:"notification-service"})

(booking-service)-[:DEPENDS_ON]->(payment-service)
(payment-service)-[:DEPENDS_ON]->(postgres)
(payment-service)-[:DEPENDS_ON]->(kafka)
(notification-service)-[:DEPENDS_ON]->(kafka)
