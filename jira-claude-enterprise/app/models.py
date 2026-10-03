from datetime import datetime, timezone
from sqlalchemy import String, DateTime, JSON, Integer
from sqlalchemy.orm import Mapped, mapped_column
from app.db import Base

def utcnow():
    return datetime.now(timezone.utc)

class WorkItem(Base):
    __tablename__ = "work_items"
    jira_key: Mapped[str] = mapped_column(String(64), primary_key=True)
    repository: Mapped[str] = mapped_column(String(255))
    base_branch: Mapped[str] = mapped_column(String(255))
    feature_branch: Mapped[str | None] = mapped_column(String(255))
    approved_push_commit: Mapped[str | None] = mapped_column(String(64))
    approved_pr_commit: Mapped[str | None] = mapped_column(String(64))
    state: Mapped[str] = mapped_column(String(32))
    created_by: Mapped[str] = mapped_column(String(255))
    updated_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow, onupdate=utcnow)

class AuditEvent(Base):
    __tablename__ = "audit_events"
    id: Mapped[int] = mapped_column(Integer, primary_key=True)
    jira_key: Mapped[str] = mapped_column(String(64), index=True)
    actor: Mapped[str] = mapped_column(String(255))
    action: Mapped[str] = mapped_column(String(64))
    commit_sha: Mapped[str | None] = mapped_column(String(64))
    details: Mapped[dict] = mapped_column(JSON, default=dict)
    created_at: Mapped[datetime] = mapped_column(DateTime(timezone=True), default=utcnow)
