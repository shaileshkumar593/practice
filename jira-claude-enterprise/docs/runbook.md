# Operations Runbook

## Startup

```bash
alembic upgrade head
uvicorn app.main:app --host 0.0.0.0 --port 8000
```

## Before production

- Configure OIDC.
- Configure GitHub App.
- Store private key in secret manager.
- Configure Jira credentials.
- Configure OpenSearch.
- Configure pgvector.
- Replace stub embedding provider.
- Replace stub reranker.
- Configure OpenTelemetry.
- Configure network egress policy.
- Enable GitHub branch protection.
- Run SAST and dependency scanning.

## Incident response

If unauthorized activity is suspected:

1. Disable the affected OIDC principal.
2. Suspend the GitHub App installation.
3. Rotate Jira credentials if exposed.
4. Preserve audit events.
5. Review approval SHA and GitHub events.
6. Revoke active sessions/tokens.
7. Rebuild the affected worker image from a trusted commit.
