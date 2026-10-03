from app.retrieval import Chunk, HybridRetriever

def test_hybrid_retrieval():
    chunks = [
        Chunk("app/payment.py", "def charge_card(): payment authorization"),
        Chunk("app/user.py", "def create_user(): user registration"),
    ]
    result = HybridRetriever(chunks).search("payment charge", 1)
    assert result[0].chunk.path == "app/payment.py"
