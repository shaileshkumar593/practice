from alembic import op
import sqlalchemy as sa

revision = "0001_initial"
down_revision = None
branch_labels = None
depends_on = None

def upgrade():
    op.create_table(
        "work_items",
        sa.Column("jira_key", sa.String(64), primary_key=True),
        sa.Column("repository", sa.String(255), nullable=False),
        sa.Column("base_branch", sa.String(255), nullable=False),
        sa.Column("feature_branch", sa.String(255)),
        sa.Column("approved_push_commit", sa.String(64)),
        sa.Column("approved_pr_commit", sa.String(64)),
        sa.Column("state", sa.String(32), nullable=False),
        sa.Column("created_by", sa.String(255), nullable=False),
        sa.Column("updated_at", sa.DateTime(timezone=True), nullable=False),
    )
    op.create_table(
        "audit_events",
        sa.Column("id", sa.Integer, primary_key=True),
        sa.Column("jira_key", sa.String(64), nullable=False),
        sa.Column("actor", sa.String(255), nullable=False),
        sa.Column("action", sa.String(64), nullable=False),
        sa.Column("commit_sha", sa.String(64)),
        sa.Column("details", sa.JSON(), nullable=False),
        sa.Column("created_at", sa.DateTime(timezone=True), nullable=False),
    )

def downgrade():
    op.drop_table("audit_events")
    op.drop_table("work_items")
