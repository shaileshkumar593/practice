from dataclasses import dataclass
from pathlib import Path
import ast
import re
from rank_bm25 import BM25Okapi
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.metrics.pairwise import cosine_similarity

@dataclass(frozen=True)
class Chunk:
    path: str
    text: str
    symbol: str = ""
    language: str = ""

@dataclass(frozen=True)
class Result:
    chunk: Chunk
    score: float

class RepositoryIndexer:
    EXTS = {".py",".go",".java",".js",".ts",".tsx",".jsx",".sql",".md",".yaml",".yml",".json",".tf"}
    EXCLUDE = {".git",".venv","node_modules","__pycache__",".pytest_cache",".agent_data",".claude"}

    def __init__(self, root: Path):
        self.root = root
        self.chunks: list[Chunk] = []

    def build(self) -> list[Chunk]:
        self.chunks.clear()
        for p in self.root.rglob("*"):
            if not p.is_file() or p.suffix.lower() not in self.EXTS:
                continue
            if any(part in self.EXCLUDE for part in p.parts):
                continue
            try:
                text = p.read_text(errors="ignore")
            except OSError:
                continue
            rel = str(p.relative_to(self.root))
            symbols = self._symbols(p, text)
            for symbol, start, end in symbols:
                self.chunks.append(Chunk(rel, "\n".join(text.splitlines()[max(0,start-3):end]), symbol, p.suffix))
            lines = text.splitlines()
            for i in range(0, len(lines), 100):
                self.chunks.append(Chunk(rel, "\n".join(lines[i:i+100]), "", p.suffix))
        return self.chunks

    def _symbols(self, path, text):
        out = []
        if path.suffix == ".py":
            try:
                tree = ast.parse(text)
                lines = text.splitlines()
                for n in ast.walk(tree):
                    if isinstance(n, (ast.FunctionDef, ast.AsyncFunctionDef, ast.ClassDef)):
                        out.append((n.name, n.lineno-1, getattr(n, "end_lineno", n.lineno)))
            except SyntaxError:
                pass
        else:
            patterns = [
                r"\b(?:func|function|class|def|interface|struct)\s+([A-Za-z_][\w]*)"
            ]
            for pat in patterns:
                for m in re.finditer(pat, text):
                    out.append((m.group(1), text[:m.start()].count("\n"), text[:m.start()].count("\n")+30))
        return out

class HybridRetriever:
    def __init__(self, chunks: list[Chunk]):
        self.chunks = chunks
        self.docs = [c.text for c in chunks]
        self.tokens = [d.lower().split() for d in self.docs]
        self.bm25 = BM25Okapi(self.tokens) if self.tokens else None
        self.vectorizer = TfidfVectorizer(max_features=20000) if self.docs else None
        self.matrix = self.vectorizer.fit_transform(self.docs) if self.docs else None

    def search(self, query: str, limit: int = 12) -> list[Result]:
        if not self.chunks:
            return []
        qtokens = query.lower().split()
        lexical = self.bm25.get_scores(qtokens)
        emb = cosine_similarity(self.vectorizer.transform([query]), self.matrix)[0]
        max_lex = max(lexical) or 1
        max_emb = max(emb) or 1
        scored = []
        for i, c in enumerate(self.chunks):
            symbol = 1.0 if c.symbol and c.symbol.lower() in query.lower() else 0.0
            metadata = 1.0 if Path(c.path).suffix.lower() in query.lower() else 0.0
            score = .35*(lexical[i]/max_lex) + .35*(emb[i]/max_emb) + .20*symbol + .10*metadata
            scored.append(Result(c, float(score)))
        return sorted(scored, key=lambda x: x.score, reverse=True)[:limit]
