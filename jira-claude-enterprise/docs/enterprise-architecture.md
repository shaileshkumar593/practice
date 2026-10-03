# Enterprise Architecture

## 1. Control plane

```text
                    ┌───────────────────────┐
                    │       Jira Cloud      │
                    └───────────┬───────────┘
                                │
                                ▼
                     ┌─────────────────────┐
                     │     Jira MCP        │
                     │  read-only tools    │
                     └──────────┬──────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────────┐
│                    Retrieval / Context Plane                    │
│                                                                 │
│  OpenSearch BM25 ─────┐                                         │
│  pgvector embeddings ─┼─► metadata filter ─► reranker          │
│  AST / tree-sitter ───┤                                         │
│  repository metadata ─┘                                         │
└───────────────────────────────┬─────────────────────────────────┘
                                │
                                ▼
                         ┌───────────────┐
                         │  Claude Code  │
                         │ coding agent  │
                         └───────┬───────┘
                                 │
                  ┌──────────────┼──────────────┐
                  ▼              ▼              ▼
             Jira tools      Git tools      Test tools
                  │              │              │
                  └──────────────┼──────────────┘
                                 ▼
                         ┌────────────────┐
                         │ Human Review   │
                         │ + RBAC Gate    │
                         └───────┬────────┘
                                 │
                  APPROVED TO PUSH PROJ-123
                                 ▼
                       ┌──────────────────┐
                       │ GitHub App       │
                       │ installation     │
                       └────────┬─────────┘
                                ▼
                          GitHub PR
                                ▼
                         CI/CD / Security
```

## 2. Security boundaries

### Identity

OIDC validates:

- issuer
- audience
- signing key
- token expiry
- subject
- roles/groups

### Authorization

RBAC separates:

- developer
- reviewer
- release-manager
- admin

A reviewer can approve, but cannot implicitly obtain release permissions.

### Git

The GitHub App uses short-lived installation tokens instead of personal access tokens.

### Approval

Approval is tied to:

- Jira key
- feature branch
- exact commit SHA
- authenticated actor
- audit event

Changing the commit invalidates approval.

## 3. Retrieval

Production retrieval should use:

```text
Query
  │
  ├── BM25 / OpenSearch
  ├── pgvector semantic search
  ├── code-symbol search
  └── metadata filtering
          │
          ▼
     candidate set
          │
          ▼
       reranker
          │
          ▼
    context budgeter
          │
          ▼
      Claude Code
```

Never send the entire repository to the model.

## 4. Human-in-the-loop state machine

```text
IMPLEMENTING
     │
     ▼
READY_FOR_REVIEW
     │
     │ reviewer approval
     ▼
PUSH_APPROVED
     │
     │ exact approved SHA
     ▼
PUSHED
     │
     │ separate PR approval
     ▼
PR_APPROVED
     │
     ▼
PR_CREATED
```

No transition may skip a security gate.

## 5. Recommended deployment

Run:

- API in Kubernetes
- MCP server as isolated deployment
- worker jobs in isolated pods
- PostgreSQL managed service
- pgvector-enabled database
- OpenSearch managed cluster
- OpenTelemetry Collector
- Prometheus/Grafana or New Relic
- enterprise OIDC provider
- GitHub App
- secrets manager
- network policies
- egress restrictions

The coding agent should never have cluster-admin or production deployment credentials.
