import sys
import numpy as np
from dataclasses import dataclass
from common.openai_client import client, DEFAULT_MODEL, EMBEDDING_MODEL

@dataclass
class Document:
    id: str
    text: str

DOCS = [
    Document("doc-1", "A supplier reservation must be confirmed within five minutes."),
    Document("doc-2", "A supplier timeout should result in PENDING and reconciliation."),
    Document("doc-3", "Idempotency keys prevent duplicate bookings."),
    Document("doc-4", "Webhook event IDs should be used for deduplication."),
]

def embed(text):
    r = client.embeddings.create(model=EMBEDDING_MODEL, input=text)
    return np.asarray(r.data[0].embedding, dtype=np.float32)

class VectorStore:
    def __init__(self, docs):
        self.docs = docs
        self.vectors = np.vstack([embed(d.text) for d in docs])

    def search(self, query, k=3):
        q = embed(query)
        scores = self.vectors @ q / (
            np.linalg.norm(self.vectors, axis=1) * np.linalg.norm(q)
        )
        return [self.docs[i] for i in np.argsort(scores)[::-1][:k]]

def answer(question):
    retrieved = VectorStore(DOCS).search(question)
    context = "\n".join(f"[{d.id}] {d.text}" for d in retrieved)
    r = client.responses.create(
        model=DEFAULT_MODEL,
        input=f"Answer only from context and cite IDs.\nContext:\n{context}\nQuestion: {question}",
    )
    return r.output_text

if __name__ == "__main__":
    q = sys.argv[1] if len(sys.argv) > 1 else "What happens on supplier timeout?"
    print(answer(q))


"""
    Purpose: Give an LLM relevant external knowledge before generating an answer.

        Basic architecture:

        Documents
            ↓
        Chunking
            ↓
        Embeddings
            ↓
        Vector / Hybrid Index
            ↓
            ┌──────────────┐
        Query →│   Retrieval  │
            └──────┬───────┘
                    ↓
                Reranking
                    ↓
                Relevant Context
                    ↓
                    LLM
                    ↓
                Grounded Answer
        RAG is useful when information:

        Changes frequently

        Is private/internal

        Is too large to put directly into the prompt

        Isn't reliably contained in the model's training data

        Example:

        Company Documentation
                ↓
            Elasticsearch
                ↓
        Relevant policies
                ↓
            Groq LLM
                ↓
        "According to company policy..."

"""