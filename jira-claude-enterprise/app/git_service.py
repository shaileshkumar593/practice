import re
import subprocess
from pathlib import Path


class GitSafetyError(RuntimeError):
    pass


class GitService:
    SAFE_BRANCH = re.compile(
        r"^feature/[A-Z][A-Z0-9]+-\d+-[a-z0-9][a-z0-9._-]{0,80}$"
    )

    def __init__(self, root: str | Path, base_branch: str):
        self.root = Path(root)
        self.base_branch = base_branch

    def run(self, *args: str) -> str:
        process = subprocess.run(
            ["git", *args],
            cwd=self.root,
            text=True,
            capture_output=True,
            check=False,
            timeout=60,
        )
        if process.returncode:
            raise GitSafetyError(
                process.stderr.strip() or "git command failed"
            )
        return process.stdout.strip()

    def branch(self, jira_key: str, slug: str) -> str:
        branch = f"feature/{jira_key}-{slug}"
        if not self.SAFE_BRANCH.fullmatch(branch):
            raise GitSafetyError("Unsafe feature branch name")

        if self.run("status", "--porcelain"):
            raise GitSafetyError("Working tree must be clean")

        self.run("fetch", "--prune", "origin")
        self.run("checkout", self.base_branch)
        self.run("reset", "--hard", f"origin/{self.base_branch}")
        self.run("checkout", "-B", branch)
        return branch

    def current_branch(self) -> str:
        return self.run("branch", "--show-current")

    def head(self) -> str:
        return self.run("rev-parse", "HEAD")

    def status(self) -> str:
        return self.run("status", "--short")

    def diff(self) -> str:
        return self.run(
            "diff",
            "--stat",
            f"origin/{self.base_branch}...HEAD",
        )

    def push(self, expected_branch: str, expected_sha: str):
        if self.current_branch() != expected_branch:
            raise GitSafetyError("Branch changed after approval")

        if self.head() != expected_sha:
            raise GitSafetyError(
                "Commit changed after approval; human review required again"
            )

        if self.status():
            raise GitSafetyError(
                "Working tree changed after approval"
            )

        if expected_branch in {self.base_branch, "master"}:
            raise GitSafetyError("Protected branch")

        self.run(
            "push",
            "--set-upstream",
            "origin",
            expected_branch,
        )
