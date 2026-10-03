# Jira → MCP → Hybrid RAG → Claude Code → Human Gate → GitHub PR

FastAPI orchestration project implementing the requested architecture.

## Architecture

Jira
 |
 v
Jira MCP
 |
 v
Hybrid Retrieval
 - BM25 / keyword
 - embeddings
 - code-symbol retrieval
 - metadata filtering
 - reranking
 |
 v
Claude Code
 |
 +-- Jira Tool
 +-- Git Tool
 +-- Test Tool
 |
 v
Human Review
 |
 v
Approval Gate
 |
 v
GitHub PR
 |
 v
CI/CD

## Setup

```bash
python -m venv .venv
# Linux/macOS:
source .venv/bin/activate
# Windows:
# .venv\Scripts\Activate.ps1

pip install -r requirements.txt
cp .env.example .env
uvicorn app.main:app --reload --port 8000
```

Swagger: http://localhost:8000/docs

## Environment

```dotenv
API_KEY=change-me
JIRA_BASE_URL=https://your-company.atlassian.net
JIRA_EMAIL=developer@example.com
JIRA_API_TOKEN=...
WORKSPACE_ROOT=/absolute/path/to/target/repository
GIT_REMOTE=origin
GIT_BASE_BRANCH=main
GITHUB_REPO=owner/repository
LEXICAL_WEIGHT=0.35
EMBEDDING_WEIGHT=0.35
SYMBOL_WEIGHT=0.20
METADATA_WEIGHT=0.10
TOP_K=20
```

## Workflow

### Start

```http
POST /api/v1/work-items/PROJ-123/start
```

This:
1. Reads Jira through the Jira MCP adapter.
2. Creates a feature branch.
3. Indexes the repository.
4. Performs BM25 + embedding + symbol retrieval.
5. Applies metadata filtering.
6. Reranks results.
7. Writes `.claude/context/PROJ-123.md`.

### Claude Code

Tell Claude:

```text
Implement PROJ-123.

Read .claude/context/PROJ-123.md.
Follow CLAUDE.md.
Inspect the repository.
Implement the ticket.
Run tests and validation.
Do NOT push or create a PR.
Stop for human review.
```

### Ready for review

```http
POST /api/v1/work-items/PROJ-123/ready-for-review
```

### Human approval

```http
POST /api/v1/work-items/PROJ-123/approve
```

```json
{
  "phrase": "APPROVED TO PUSH PROJ-123"
}
```

### Push

```http
POST /api/v1/work-items/PROJ-123/push
```

The server rejects push without matching human approval.

### Create PR

```http
POST /api/v1/work-items/PROJ-123/create-pr
```

Requires the branch to already be pushed.

## Production recommendations

Replace the local TF-IDF embedding implementation with a real embedding model/vector store such as pgvector/Qdrant/Weaviate, and use the real Jira MCP server available in your environment.

Put FastAPI behind OIDC/SSO and RBAC. Use secret management for Jira/GitHub credentials. Protect main/master and require CI + human merge approval.
