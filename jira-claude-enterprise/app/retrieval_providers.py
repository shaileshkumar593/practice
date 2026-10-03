from dataclasses import dataclass
from typing import Protocol

import httpx

from app.config import settings


class EmbeddingProvider(Protocol):
    def embed(self, texts: list[str]) -> list[list[float]]: ...


class Reranker(Protocol):
    def rerank(self, query: str, documents: list[str]) -> list[float]: ...


@dataclass
class StubEmbeddingProvider:
    """Deterministic development adapter. Replace with an approved embedding provider."""

    dimensions: int = 64

    def embed(self, texts: list[str]) -> list[list[float]]:
        vectors = []
        for text in texts:
            vector = [0.0] * self.dimensions
            for token in text.lower().split():
                vector[hash(token) % self.dimensions] += 1.0
            norm = sum(v * v for v in vector) ** 0.5 or 1.0
            vectors.append([v / norm for v in vector])
        return vectors


@dataclass
class StubReranker:
    def rerank(self, query: str, documents: list[str]) -> list[float]:
        q = set(query.lower().split())
        return [
            len(q.intersection(set(doc.lower().split()))) / max(len(q), 1)
            for doc in documents
        ]


class OpenSearchRetriever:
    """
    Production retrieval boundary.

    OpenSearch should hold BM25 text fields and vector fields. The actual index
    mapping and authentication are environment-specific and intentionally kept
    outside application source.
    """

    def __init__(self, base_url: str, index: str):
        self.base_url = base_url.rstrip("/")
        self.index = index

    def search(self, query: str, limit: int = 12) -> list[dict]:
        body = {
            "size": limit,
            "query": {
                "multi_match": {
                    "query": query,
                    "fields": ["text^2", "symbol", "path"],
                }
            },
        }
        with httpx.Client(timeout=15) as client:
            response = client.post(
                f"{self.base_url}/{self.index}/_search",
                json=body,
            )
            response.raise_for_status()
            return response.json().get("hits", {}).get("hits", [])


def embedding_provider() -> EmbeddingProvider:
    # Enterprise deployments can inject an OpenAI/Azure/Bedrock implementation.
    return StubEmbeddingProvider()


def reranker() -> Reranker:
    return StubReranker()
