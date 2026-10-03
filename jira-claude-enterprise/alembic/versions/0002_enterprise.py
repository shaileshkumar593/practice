from alembic import op
import sqlalchemy as sa

revision = "0002_enterprise"
down_revision = "0001_initial"
branch_labels = None
depends_on = None


def upgrade():
    op.create_table(
        "code_embeddings",
        sa.Column("id", sa.Integer, primary_key=True),
        sa.Column("repository", sa.String(255), nullable=False),
        sa.Column("branch", sa.String(255), nullable=False),
        sa.Column("path", sa.String(1024), nullable=False),
        sa.Column("symbol", sa.String(512)),
        sa.Column("content", sa.Text, nullable=False),
        sa.Column("embedding_json", sa.Text, nullable=False),
    )
    op.create_index(
        "ix_code_embeddings_repository",
        "code_embeddings",
        ["repository"],
    )
    op.create_index(
        "ix_code_embeddings_path",
        "code_embeddings",
        ["path"],
    )


def downgrade():
    op.drop_index(
        "ix_code_embeddings_path",
        table_name="code_embeddings",
    )
    op.drop_index(
        "ix_code_embeddings_repository",
        table_name="code_embeddings",
    )
    op.drop_table("code_embeddings")
