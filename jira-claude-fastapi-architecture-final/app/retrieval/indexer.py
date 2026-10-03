from pathlib import Path
import ast
import re
from .models import CodeChunk

EXT = {
    ".py": "python", ".go": "go", ".java": "java",
    ".js": "javascript", ".jsx": "javascript",
    ".ts": "typescript", ".tsx": "typescript",
    ".sql": "sql", ".md": "markdown",
    ".yaml": "yaml", ".yml": "yaml", ".json": "json",
    ".tf": "terraform",
}
EXCLUDED = {".git", ".venv", "venv", "node_modules", "__pycache__",
            ".pytest_cache", "dist", "build", ".idea", ".vscode",
            ".agent_data", ".claude"}

class RepositoryIndexer:
    def index(self, root: Path):
        out = []
        for path in root.rglob("*"):
            if not path.is_file() or any(x in EXCLUDED for x in path.parts):
                continue
            language = EXT.get(path.suffix.lower())
            if not language:
                continue
            try:
                text = path.read_text(encoding="utf-8", errors="ignore")
            except OSError:
                continue
            out.extend(self._file(path, text, language))
        return out

    def _file(self, path, text, language):
        lines = text.splitlines()
        result = []
        for symbol, start, end in extract_symbols(text, language):
            result.append(CodeChunk(
                path=str(path.relative_to(Path.cwd())),
                text="\n".join(lines[start-1:end])[:16000],
                language=language,
                symbol=symbol,
                start_line=start,
                end_line=end,
                chunk_type="symbol",
            ))
        for start in range(0, len(lines), 100):
            end = min(start + 120, len(lines))
            piece = "\n".join(lines[start:end])
            if piece.strip():
                result.append(CodeChunk(
                    path=str(path.relative_to(Path.cwd())),
                    text=piece[:16000],
                    language=language,
                    symbol=None,
                    start_line=start+1,
                    end_line=end,
                ))
        return result

def extract_symbols(text, language):
    lines = text.splitlines()
    result = []

    if language == "python":
        try:
            tree = ast.parse(text)
            for node in ast.walk(tree):
                if isinstance(node, (ast.FunctionDef, ast.AsyncFunctionDef, ast.ClassDef)):
                    result.append((node.name, node.lineno, getattr(node, "end_lineno", node.lineno)))
        except SyntaxError:
            pass

    elif language in {"go", "java", "javascript", "typescript"}:
        patterns = [
            r"^\s*(?:func|function|class|interface|type)\s+([A-Za-z_]\w*)",
            r"^\s*(?:export\s+)?(?:async\s+)?function\s+([A-Za-z_]\w*)",
        ]
        for i, line in enumerate(lines, 1):
            for pattern in patterns:
                m = re.match(pattern, line)
                if m:
                    result.append((m.group(1), i, min(i + 80, len(lines))))
                    break
    return result
