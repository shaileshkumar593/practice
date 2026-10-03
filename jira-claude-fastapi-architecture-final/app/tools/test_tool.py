import subprocess
from app.config import settings

class TestTool:
    COMMANDS = {
        "python": ["pytest", "-q"],
        "go": ["go", "test", "./..."],
        "node": ["npm", "test"],
    }

    def run(self, ecosystem="python"):
        command = self.COMMANDS.get(ecosystem)
        if not command:
            raise ValueError("Unsupported ecosystem")
        p = subprocess.run(
            command,
            cwd=settings.workspace_root,
            capture_output=True,
            text=True,
        )
        return {
            "command": " ".join(command),
            "return_code": p.returncode,
            "passed": p.returncode == 0,
            "stdout": p.stdout[-20000:],
            "stderr": p.stderr[-20000:],
        }
