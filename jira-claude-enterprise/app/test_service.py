import subprocess
from pathlib import Path

ALLOWED = {
    "python": ["pytest", "-q"],
    "go": ["go", "test", "./..."],
    "node": ["npm", "test"],
}

class TestService:
    def __init__(self, root: Path): self.root = root

    def run(self, ecosystem: str = "python"):
        cmd = ALLOWED.get(ecosystem)
        if not cmd:
            raise ValueError("Unsupported test ecosystem")
        p = subprocess.run(cmd, cwd=self.root, text=True, capture_output=True,
                           timeout=300, check=False)
        return {"command": cmd, "exit_code": p.returncode,
                "stdout": p.stdout[-12000:], "stderr": p.stderr[-12000:]}
