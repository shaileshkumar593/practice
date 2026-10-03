from dataclasses import dataclass, field

@dataclass
class CodeChunk:
    path: str
    text: str
    language: str
    symbol: str | None
    start_line: int
    end_line: int
    chunk_type: str = "code"

@dataclass
class RetrievalResult:
    chunk: CodeChunk
    lexical_score: float
    embedding_score: float
    symbol_score: float
    metadata_score: float
    final_score: float
    reasons: list[str] = field(default_factory=list)
