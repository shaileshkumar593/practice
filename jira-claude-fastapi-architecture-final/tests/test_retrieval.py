from app.retrieval.models import CodeChunk
from app.retrieval.hybrid import HybridRetriever

def test_hybrid_retrieval():
    chunks = [
        CodeChunk("user.py", "def create_user(user): return user", "python", "create_user", 1, 1, "symbol"),
        CodeChunk("payment.py", "def charge_card(card): return card", "python", "charge_card", 1, 1, "symbol"),
    ]
    results = HybridRetriever(chunks).search("create user", top_k=1)
    assert results[0].chunk.symbol == "create_user"
