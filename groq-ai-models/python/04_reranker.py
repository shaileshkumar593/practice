from sentence_transformers import CrossEncoder

model = CrossEncoder("cross-encoder/ms-marco-MiniLM-L-6-v2")
query = "How do I make a payment API idempotent?"
docs = [
    "Use an Idempotency-Key and persist the request/result with a unique constraint.",
    "Kubernetes schedules containers on nodes.",
    "Redis can be used for caching.",
    "Payment gateways may support provider-side idempotency keys.",
]
scores = model.predict([(query,d) for d in docs])
for score, doc in sorted(zip(scores,docs), reverse=True):
    print(f"{float(score):.4f} {doc}")
