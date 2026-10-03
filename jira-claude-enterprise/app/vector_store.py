from sqlalchemy import Text, String
from sqlalchemy.orm import Mapped, mapped_column

from app.db import Base


class CodeEmbedding(Base):
    """
    pgvector-ready logical model.

    Add a native Vector(N) column in the production migration after choosing the
    organization's embedding dimensions/provider.
    """

    __tablename__ = "code_embeddings"

    id: Mapped[int] = mapped_column(primary_key=True)
    repository: Mapped[str] = mapped_column(String(255), index=True)
    branch: Mapped[str] = mapped_column(String(255), index=True)
    path: Mapped[str] = mapped_column(String(1024), index=True)
    symbol: Mapped[str | None] = mapped_column(String(512))
    content: Mapped[str] = mapped_column(Text)
    embedding_json: Mapped[str] = mapped_column(Text)
