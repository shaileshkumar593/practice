from common.clients import chat,embed
from sentence_transformers import CrossEncoder
import numpy as np

docs=[
 "Idempotency keys prevent duplicate payment requests.",
 "Use a unique database constraint for an idempotency key.",
 "A payment timeout does not prove failure; query provider status or wait for webhook.",
 "Webhook event IDs should be persisted to prevent duplicate processing.",
 "Redis is useful for caching but should not be the only source of payment truth."
]
query="What should happen when a payment request times out?"
def cos(a,b):
    a,b=np.array(a),np.array(b)
    return float(np.dot(a,b)/(np.linalg.norm(a)*np.linalg.norm(b)))

vs=embed(docs); q=embed([query])[0]
cands=[d for _,d in sorted([(cos(q,v),d) for v,d in zip(vs,docs)],reverse=True)[:4]]
rr=CrossEncoder("cross-encoder/ms-marco-MiniLM-L-6-v2")
scores=rr.predict([(query,d) for d in cands])
ctx=[d for _,d in sorted(zip(scores,cands),reverse=True)]
print(chat("Context:\n"+"\n".join(ctx)+"\n\nQuestion:"+query,
           "Answer only from the provided context."))
