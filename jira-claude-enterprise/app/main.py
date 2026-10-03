from fastapi import Body, Depends, FastAPI
from prometheus_client import Counter, generate_latest
from sqlalchemy.orm import Session
from starlette.responses import Response

from app.auth import Principal, current_principal, require_permission
from app.db import get_db
from app.service import Orchestrator
from app.workflow import Workflow
from app.git_service import GitService
from app.config import settings

app = FastAPI(
    title="Jira → Claude Enterprise Orchestrator",
    version="2.0.0",
    docs_url="/docs" if settings.app_env != "production" else None,
)

requests = Counter(
    "orchestrator_requests_total",
    "Orchestrator API requests",
    ["route"],
)


@app.get("/health")
def health():
    return {"status": "ok", "service": "jira-claude-orchestrator"}


@app.get("/metrics")
def metrics():
    return Response(
        generate_latest(),
        media_type="text/plain; version=0.0.4",
    )


@app.post("/api/v1/work-items/{jira_key}/start")
async def start(
    jira_key: str,
    db: Session = Depends(get_db),
    principal: Principal = Depends(current_principal),
):
    require_permission(principal, "workitem:start")
    requests.labels("start").inc()
    return await Orchestrator(db, principal.subject).start(jira_key)


@app.get("/api/v1/work-items/{jira_key}")
def inspect(
    jira_key: str,
    db: Session = Depends(get_db),
    principal: Principal = Depends(current_principal),
):
    require_permission(principal, "workitem:inspect")
    requests.labels("inspect").inc()
    return Orchestrator(db, principal.subject).inspect(jira_key)


@app.post("/api/v1/work-items/{jira_key}/ready-for-review")
def ready(
    jira_key: str,
    db: Session = Depends(get_db),
    principal: Principal = Depends(current_principal),
):
    require_permission(principal, "workitem:ready")
    return Orchestrator(db, principal.subject).ready(jira_key)


@app.post("/api/v1/work-items/{jira_key}/approve-push")
def approve_push(
    jira_key: str,
    phrase: str = Body(..., embed=True),
    db: Session = Depends(get_db),
    principal: Principal = Depends(current_principal),
):
    require_permission(principal, "workitem:approve_push")
    git = GitService(
        settings.workspace_root,
        settings.github_base_branch,
    )
    return Workflow(db, principal.subject, git).approve_push(jira_key, phrase)


@app.post("/api/v1/work-items/{jira_key}/push")
def push(
    jira_key: str,
    db: Session = Depends(get_db),
    principal: Principal = Depends(current_principal),
):
    require_permission(principal, "workitem:push")
    git = GitService(
        settings.workspace_root,
        settings.github_base_branch,
    )
    return Workflow(db, principal.subject, git).push(jira_key)


@app.post("/api/v1/work-items/{jira_key}/approve-pr")
def approve_pr(
    jira_key: str,
    phrase: str = Body(..., embed=True),
    db: Session = Depends(get_db),
    principal: Principal = Depends(current_principal),
):
    require_permission(principal, "workitem:approve_pr")
    git = GitService(
        settings.workspace_root,
        settings.github_base_branch,
    )
    return Workflow(db, principal.subject, git).approve_pr(jira_key, phrase)


@app.post("/api/v1/work-items/{jira_key}/create-pr")
def create_pr(
    jira_key: str,
    db: Session = Depends(get_db),
    principal: Principal = Depends(current_principal),
):
    require_permission(principal, "workitem:push")
    git = GitService(
        settings.workspace_root,
        settings.github_base_branch,
    )
    return Workflow(db, principal.subject, git).create_pr(jira_key)
