import numpy as np
from common.openai_client import client, EMBEDDING_MODEL

def embed(text: str):
    r = client.embeddings.create(model=EMBEDDING_MODEL, input=text)
    return r.data[0].embedding

def cosine(a, b):
    a, b = np.asarray(a), np.asarray(b)
    d = np.linalg.norm(a) * np.linalg.norm(b)
    return float(np.dot(a, b) / d) if d else 0.0

if __name__ == "__main__":
    q = embed("How do I prevent duplicate payment?")
    docs = [
        "Use an idempotency key with a unique database constraint.",
        "Use Redis for product caching.",
        "Use a CDN for static images.",
    ]
    for doc in sorted(docs, key=lambda x: cosine(q, embed(x)), reverse=True):
        print(doc)



"""
    Purpose: Convert text, images, or other content into numerical vectors.

    Example:

    "How do I reset my password?"
                ↓
        Embedding Model
                ↓
    [0.12, -0.45, 0.78, ...]
    Similar meanings produce vectors that are closer together.

    Used for:

    Semantic search

    RAG

    Recommendation systems

    Clustering

    Duplicate detection

    Classification

    Similarity is commonly measured using cosine similarity.

    Embeddings are stored in a vector database or vector-capable search engine.

    Examples: OpenAI embedding models, BGE, E5.

    Interview point:
    An embedding model doesn't normally generate the final answer. It converts information into a representation useful for retrieval or similarity.
"""
