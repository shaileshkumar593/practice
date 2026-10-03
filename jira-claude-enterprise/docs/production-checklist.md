# Production Checklist

## Identity and authorization
- Enterprise OIDC/JWKS validation enabled.
- RBAC: developer, reviewer, release-manager.
- Reviewer cannot approve their own change where separation of duties is required.
- GitHub App credentials scoped to the target repository.

## Repository safety
- Workspace isolated per job/container.
- No host Docker socket mounted.
- No arbitrary shell execution from Jira or model output.
- Protected branches enforced on GitHub.
- Push permission limited to feature branches.
- Approval bound to exact commit SHA.

## Retrieval
- BM25 in OpenSearch/Elasticsearch.
- Embeddings in pgvector/Qdrant/Weaviate.
- AST/tree-sitter symbol index.
- Metadata filtering by repository, branch, language and ownership.
- Reranker after candidate retrieval.
- Retrieval traces retained for debugging.

## State and audit
- PostgreSQL with migrations.
- Immutable audit events shipped to SIEM.
- Idempotency keys for start/approval/push operations.
- Approval TTL and revocation support.
- Correlation/request IDs.

## Claude Code
- Use a dedicated service account.
- Restrict filesystem to workspace.
- Restrict network egress.
- Read-only Jira context where possible.
- Explicit human gate before irreversible operations.

## CI/CD
- Required checks before PR.
- SAST, dependency scanning, secret scanning.
- Unit/integration tests.
- Build and artifact signing.
- Deployment remains outside the coding agent's default permissions.
