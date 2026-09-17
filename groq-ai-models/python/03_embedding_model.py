from common.clients import embed
import numpy as np

docs = [
    "PostgreSQL supports ACID transactions.",
    "Elasticsearch is commonly used for search.",
    "Kafka is a distributed event streaming platform.",
]
vectors = embed(docs)
q = embed(["transactional relational database"])[0]

def cosine(a,b):
    a,b=np.array(a),np.array(b)
    return float(np.dot(a,b)/(np.linalg.norm(a)*np.linalg.norm(b)))

for score, doc in sorted([(cosine(q,v),d) for v,d in zip(vectors,docs)], reverse=True):
    print(round(score,4), doc)
