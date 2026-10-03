from rank_bm25 import BM25Okapi
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.metrics.pairwise import cosine_similarity
from app.config import settings
from .models import RetrievalResult

class HybridRetriever:
    def __init__(self, chunks):
        self.chunks = chunks
        self.bm25 = BM25Okapi([
            self.tokens(c.text + " " + (c.symbol or "") + " " + c.path)
            for c in chunks
        ])
        self.vectorizer = TfidfVectorizer(
            lowercase=True,
            ngram_range=(1, 2),
            max_features=60000,
            token_pattern=r"(?u)\b[\w./:-]{2,}\b",
        )
        self.matrix = self.vectorizer.fit_transform([
            f"{c.path} {c.symbol or ''} {c.text}" for c in chunks
        ])

    def search(self, query, top_k=None, language=None, path_prefix=None):
        top_k = top_k or settings.top_k
        lexical = normalize(self.bm25.get_scores(self.tokens(query)))
        q = self.vectorizer.transform([query])
        embedding = normalize(cosine_similarity(q, self.matrix).ravel())
        ql = query.lower()
        results = []

        for i, chunk in enumerate(self.chunks):
            if language and chunk.language != language:
                continue
            if path_prefix and not chunk.path.startswith(path_prefix):
                continue

            symbol = 1.0 if chunk.symbol and chunk.symbol.lower() in ql else 0.0
            metadata = 1.0 if path_prefix and chunk.path.startswith(path_prefix) else 0.0

            final = (
                settings.lexical_weight * lexical[i] +
                settings.embedding_weight * embedding[i] +
                settings.symbol_weight * symbol +
                settings.metadata_weight * metadata
            )

            reasons = []
            if lexical[i] > 0: reasons.append("bm25")
            if embedding[i] > 0: reasons.append("embedding")
            if symbol > 0: reasons.append("symbol")
            if metadata > 0: reasons.append("metadata")

            results.append(RetrievalResult(
                chunk=chunk,
                lexical_score=float(lexical[i]),
                embedding_score=float(embedding[i]),
                symbol_score=symbol,
                metadata_score=metadata,
                final_score=float(final),
                reasons=reasons,
            ))

        results.sort(key=lambda x: x.final_score, reverse=True)
        return results[:top_k]

    @staticmethod
    def tokens(text):
        return text.lower().split()

def normalize(values):
    if len(values) == 0:
        return values
    lo, hi = float(min(values)), float(max(values))
    if hi == lo:
        return [0.0 for _ in values]
    return [(float(v)-lo)/(hi-lo) for v in values]
