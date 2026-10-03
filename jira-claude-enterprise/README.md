# Jira → Claude Code Production Orchestrator

Production-oriented reference implementation for a strict human-in-the-loop workflow:

Jira → MCP adapter → Hybrid Retrieval → Claude Code → deterministic tools
→ tests → human review gate → Git push → GitHub PR → CI/CD

## Security model

The service is deliberately conservative:

- Never writes to `main`/`master`.
- Feature branches are created from a freshly fetched protected base.
- `git push` requires an explicit approval token:
  `APPROVED TO PUSH PROJ-123`
- Approval is bound to the Jira key, repository, branch and current commit.
- Approval expires after a configurable TTL.
- Any working-tree change after approval invalidates approval.
- PR creation requires a pushed commit and a second explicit approval:
  `APPROVED TO CREATE PR PROJ-123`
- The service never executes arbitrary shell commands supplied by Jira/LLM.
- Test commands come only from a configured allow-list.
- API authentication uses OIDC/JWT in production mode.
- State is persisted in PostgreSQL.
- Secrets are injected through environment/secret manager, never committed.

> This repository is a production-oriented reference architecture. Replace the example identity provider and Jira/GitHub credentials with your organization's controls before deployment.

## Components

- FastAPI API
- PostgreSQL + SQLAlchemy 2
- Jira REST adapter behind an MCP-compatible service boundary
- Hybrid retrieval interfaces: BM25 + embeddings + code symbols + metadata
- pgvector-ready embedding store boundary
- Git safety service
- deterministic test runner
- immutable audit events
- approval state machine
- Prometheus metrics
- OpenTelemetry hooks
- Docker Compose for local integration
- GitHub Actions CI

## Workflow

1. `POST /api/v1/work-items/PROJ-123/start`
2. Service fetches Jira context and creates a feature branch.
3. Retrieval builds `.claude/context/PROJ-123.md`.
4. Claude Code works only from that context and repository.
5. Developer runs tests and marks work ready.
6. Reviewer inspects diff/status and calls approve with exact phrase.
7. Push is allowed only for the approved commit.
8. PR creation requires a separate approval.

See `docs/production-checklist.md`.

## Run locally

```bash
cp .env.example .env
docker compose up --build
```

API:
`http://localhost:8000/docs`

Run tests:

```bash
docker compose exec api pytest -q
```

## Important production substitutions

The included retrieval engine has provider interfaces. For production, connect:

- embeddings: OpenAI/Azure OpenAI or your approved embedding provider
- vector DB: pgvector/Qdrant/Weaviate
- lexical: OpenSearch/Elasticsearch BM25
- reranker: approved cross-encoder/reranking API
- secrets: AWS Secrets Manager / GCP Secret Manager / Vault
- identity: enterprise OIDC
- telemetry: OpenTelemetry collector + New Relic/Grafana/Prometheus
- Git hosting: GitHub App with least-privilege permissions

Do not give the API process a personal developer PAT.
