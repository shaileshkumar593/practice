import os
from dotenv import load_dotenv
from sentence_transformers import CrossEncoder

load_dotenv()
model = CrossEncoder(os.getenv(
    "RERANKER_MODEL", "cross-encoder/ms-marco-MiniLM-L-6-v2"
))

def rerank(query, documents, top_k=3):
    scores = model.predict([(query, d) for d in documents])
    rows = sorted(
        zip(documents, scores), key=lambda x: float(x[1]), reverse=True
    )
    return [{"document": d, "score": float(s)} for d, s in rows[:top_k]]

if __name__ == "__main__":
    q = "How can a payment API prevent duplicate charges?"
    docs = [
        "Idempotency keys prevent duplicate payment processing.",
        "PostgreSQL indexes speed up queries.",
        "CDNs cache static assets.",
        "Payment webhooks confirm asynchronous payments.",
    ]
    for x in rerank(q, docs):
        print(x)


"""
    Purpose: Re-rank documents retrieved by a search system according to their relevance to the query.

    Typical pipeline:

        Query
        ↓
        Vector/BM25 Search
        ↓
        Top 50 Documents
        ↓
        Reranker
        ↓
        Top 5 Relevant Documents
    Initial retrieval is optimized for speed.

    Reranking is optimized for relevance.

    A common approach is a cross-encoder, which evaluates:

    (query, document)
    together.

    Especially useful in RAG and enterprise search.

    It can substantially improve the quality of the context given to an LLM.

    Interview point:
    Don't use a reranker across millions of documents. First retrieve a manageable candidate set, then rerank it.

"""
