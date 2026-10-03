from datetime import datetime, timezone, timedelta

from fastapi import HTTPException
from sqlalchemy.orm import Session

from app.config import settings
from app.models import AuditEvent, WorkItem
from app.github_app import GitHubAppClient
from app.git_service import GitService


class Workflow:
    def __init__(self, db: Session, actor: str, git: GitService):
        self.db = db
        self.actor = actor
        self.git = git

    def item(self, key: str) -> WorkItem:
        item = self.db.get(WorkItem, key)
        if not item:
            raise HTTPException(404, "Unknown work item")
        return item

    def audit(self, key: str, action: str, sha: str | None = None, **details):
        self.db.add(
            AuditEvent(
                jira_key=key,
                actor=self.actor,
                action=action,
                commit_sha=sha,
                details=details,
            )
        )
        self.db.commit()

    def approve_push(self, key: str, phrase: str):
        item = self.item(key)
        if phrase != f"APPROVED TO PUSH {key}":
            raise HTTPException(400, "Exact push approval phrase required")
        if item.state != "READY_FOR_REVIEW":
            raise HTTPException(409, "Work item is not ready for review")

        sha = self.git.head()
        item.approved_push_commit = sha
        item.state = "PUSH_APPROVED"
        item.updated_at = datetime.now(timezone.utc)
        self.db.commit()

        self.audit(
            key,
            "PUSH_APPROVED",
            sha,
            expires_at=(
                datetime.now(timezone.utc)
                + timedelta(seconds=settings.approval_ttl_seconds)
            ).isoformat(),
        )
        return {"state": item.state, "approved_commit": sha}

    def push(self, key: str):
        item = self.item(key)

        if item.state != "PUSH_APPROVED":
            raise HTTPException(403, "Push has not been approved")

        if self.git.current_branch() != item.feature_branch:
            raise HTTPException(409, "Feature branch changed")

        if self.git.status():
            raise HTTPException(409, "Working tree changed after approval")

        current_sha = self.git.head()
        if current_sha != item.approved_push_commit:
            raise HTTPException(409, "Approved commit changed")

        self.git.push(item.feature_branch, current_sha)

        item.state = "PUSHED"
        item.updated_at = datetime.now(timezone.utc)
        self.db.commit()

        self.audit(key, "PUSHED", current_sha)
        return {"state": item.state, "commit": current_sha}

    def approve_pr(self, key: str, phrase: str):
        item = self.item(key)

        if phrase != f"APPROVED TO CREATE PR {key}":
            raise HTTPException(400, "Exact PR approval phrase required")

        if item.state != "PUSHED":
            raise HTTPException(409, "Branch must be pushed before PR approval")

        sha = self.git.head()
        item.approved_pr_commit = sha
        item.state = "PR_APPROVED"
        item.updated_at = datetime.now(timezone.utc)
        self.db.commit()

        self.audit(key, "PR_APPROVED", sha)
        return {"state": item.state, "approved_commit": sha}

    def create_pr(self, key: str):
        item = self.item(key)

        if item.state != "PR_APPROVED":
            raise HTTPException(403, "PR creation has not been approved")

        sha = self.git.head()
        if sha != item.approved_pr_commit:
            raise HTTPException(409, "PR approval no longer matches HEAD")

        result = GitHubAppClient().create_pull_request(
            head=item.feature_branch,
            base=item.base_branch,
            title=f"{key}: implementation",
            body=f"Automated implementation for {key}.\n\nReviewed commit: {sha}",
        )

        item.state = "PR_CREATED"
        item.updated_at = datetime.now(timezone.utc)
        self.db.commit()

        self.audit(key, "PR_CREATED", sha, pr_url=result.get("html_url"))
        return {"state": item.state, "url": result.get("html_url")}
