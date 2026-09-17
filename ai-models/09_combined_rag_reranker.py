import os
import numpy as np
from dotenv import load_dotenv
from openai import OpenAI
from sentence_transformers import CrossEncoder

load_dotenv()
client = OpenAI()
EMBED = os.getenv("OPENAI_EMBEDDING_MODEL", "text-embedding-3-small")
GEN = os.getenv("OPENAI_MODEL", "gpt-5.6-luna")
RERANK = os.getenv("RERANKER_MODEL", "cross-encoder/ms-marco-MiniLM-L-6-v2")

docs = [
    ("1", "Idempotency keys prevent duplicate payment charges."),
    ("2", "Unknown payment outcomes are marked PENDING and reconciled by webhook or status API."),
    ("3", "PostgreSQL transactions provide atomicity."),
    ("4", "Webhook events should be deduplicated using provider event IDs."),
]

def embed(text):
    return np.asarray(
        client.embeddings.create(model=EMBED, input=text).data[0].embedding,
        dtype=np.float32,
    )

vectors = np.vstack([embed(x[1]) for x in docs])
reranker = CrossEncoder(RERANK)

def rag(question):
    q = embed(question)
    scores = vectors @ q / (np.linalg.norm(vectors, axis=1) * np.linalg.norm(q))
    candidates = [docs[i] for i in np.argsort(scores)[::-1][:4]]
    rs = reranker.predict([(question, text) for _, text in candidates])
    ranked = sorted(zip(candidates, rs), key=lambda x: float(x[1]), reverse=True)[:2]
    context = "\n".join(f"[{i}] {t}" for (i, t), _ in ranked)
    r = client.responses.create(
        model=GEN,
        input=f"Answer only from context and cite IDs.\nContext:\n{context}\nQuestion: {question}",
    )
    return r.output_text

if __name__ == "__main__":
    print(rag("How should I handle a payment timeout?"))
