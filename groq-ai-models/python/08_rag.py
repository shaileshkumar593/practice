from common.clients import chat,embed
import numpy as np

docs=[
 "Idempotency prevents duplicate payment processing. Store the key with a unique constraint.",
 "Payment webhooks must be authenticated, deduplicated, persisted and processed safely.",
 "A timeout after sending a payment request creates an unknown outcome. Reconcile before retrying.",
 "PostgreSQL transactions provide atomic payment updates and outbox events."
]
def cos(a,b):
    a,b=np.array(a),np.array(b)
    return float(np.dot(a,b)/(np.linalg.norm(a)*np.linalg.norm(b)))

query="How should I handle duplicate and timed-out payments?"
vs=embed(docs); q=embed([query])[0]
ctx=[d for _,d in sorted([(cos(q,v),d) for v,d in zip(vs,docs)],reverse=True)[:3]]

prompt="Answer ONLY from this context.\nContext:\n"+"\n".join("- "+x for x in ctx)+"\nQuestion: "+query
print(chat(prompt,"You are a grounded RAG assistant."))
