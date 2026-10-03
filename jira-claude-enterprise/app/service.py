from pathlib import Path
from datetime import datetime, timezone, timedelta
from sqlalchemy.orm import Session
from fastapi import HTTPException
from app.models import WorkItem, AuditEvent
from app.config import settings
from app.jira import JiraMCPAdapter
from app.git_service import GitService
from app.retrieval import RepositoryIndexer, HybridRetriever

class Orchestrator:
    def __init__(self, db: Session, actor: str):
        self.db, self.actor = db, actor
        self.jira = JiraMCPAdapter()
        self.git = GitService(Path(settings.workspace_root), settings.github_base_branch)

    def _audit(self, key, action, sha=None, details=None):
        self.db.add(AuditEvent(jira_key=key, actor=self.actor, action=action,
                               commit_sha=sha, details=details or {}))
        self.db.commit()

    async def start(self, key: str):
        item = self.db.get(WorkItem, key)
        if item:
            raise HTTPException(409, "Work item already exists")
        issue = await self.jira.get_issue(key)
        branch = self.git.branch(key, issue.summary.lower().replace(" ", "-")[:50])
        chunks = RepositoryIndexer(Path(settings.workspace_root)).build()
        results = HybridRetriever(chunks).search(
            f"{issue.summary}\n{issue.description}\n{issue.acceptance_criteria}",
            settings.max_retrieval_results)
        ctx = Path(settings.workspace_root) / ".claude" / "context"
        ctx.mkdir(parents=True, exist_ok=True)
        (ctx / f"{key}.md").write_text(
            f"# {key}\n\n## Summary\n{issue.summary}\n\n"
            f"## Description\n{issue.description}\n\n## Acceptance Criteria\n"
            f"{issue.acceptance_criteria}\n\n## Linked Issues\n"
            + "\n".join(issue.links) + "\n\n## Retrieved Code\n"
            + "\n\n".join(f"### {r.chunk.path}::{r.chunk.symbol}\n{r.chunk.text}" for r in results)
        )
        item = WorkItem(jira_key=key, repository=settings.github_repository,
                        base_branch=settings.github_base_branch, feature_branch=branch,
                        state="IMPLEMENTING", created_by=self.actor)
        self.db.add(item)
        self.db.commit()
        self._audit(key, "STARTED", self.git.head(), {"branch": branch})
        return self.inspect(key)

    def inspect(self, key):
        item = self.db.get(WorkItem, key)
        if not item: raise HTTPException(404, "Unknown work item")
        return {"jira_key": key, "state": item.state, "branch": item.feature_branch,
                "head": self.git.head(), "status": self.git.status(), "diff": self.git.diff()}

    def ready(self, key):
        item = self._item(key)
        if item.state != "IMPLEMENTING": raise HTTPException(409, "Invalid state")
        if self.git.current_branch() != item.feature_branch: raise HTTPException(409, "Wrong branch")
        if self.git.status(): raise HTTPException(409, "Uncommitted changes must be committed before review")
        item.state = "READY_FOR_REVIEW"
        item.updated_at = datetime.now(timezone.utc)
        self.db.commit()
        self._audit(key, "READY_FOR_REVIEW", self.git.head())
        return self.inspect(key)

    def approve_push(self, key, phrase):
        item = self._item(key)
        expected = f"APPROVED TO PUSH {key}"
        if phrase != expected: raise HTTPException(400, "Exact approval phrase required")
        if item.state != "READY_FOR_REVIEW": raise HTTPException(409, "Not ready for review")
        sha = self.git.head()
        item.approved_push_commit = sha
        item.state = "PUSH_APPROVED"
        self.db.commit()
        self._audit(key, "PUSH_APPROVED", sha)
        return self.inspect(key)

    def push(self, key):
        item = self._item(key)
        if item.state != "PUSH_APPROVED": raise HTTPException(403, "Push not approved")
        if self.git.head() != item.approved_push_commit: raise HTTPException(409, "Approval invalidated")
        self.git.push(item.feature_branch, item.approved_push_commit)
        item.state = "PUSHED"
        self.db.commit()
        self._audit(key, "PUSHED", item.approved_push_commit)
        return self.inspect(key)

    def approve_pr(self, key, phrase):
        item = self._item(key)
        expected = f"APPROVED TO CREATE PR {key}"
        if phrase != expected: raise HTTPException(400, "Exact PR approval phrase required")
        if item.state != "PUSHED": raise HTTPException(409, "Push required first")
        item.approved_pr_commit = self.git.head()
        item.state = "PR_APPROVED"
        self.db.commit()
        self._audit(key, "PR_APPROVED", item.approved_pr_commit)
        return self.inspect(key)

    def _item(self, key):
        item = self.db.get(WorkItem, key)
        if not item: raise HTTPException(404, "Unknown work item")
        return item
