import re
import subprocess
from app.config import settings

class GitTool:
    def run(self, *args):
        p = subprocess.run(
            ["git", *args],
            cwd=settings.workspace_root,
            capture_output=True,
            text=True,
        )
        if p.returncode:
            raise RuntimeError(p.stderr.strip() or "Git failed")
        return p.stdout.strip()

    def current_branch(self):
        return self.run("branch", "--show-current")

    def create_branch(self, jira_key, summary):
        slug = re.sub(r"[^a-zA-Z0-9]+", "-", summary.lower()).strip("-")[:60]
        branch = f"feature/{jira_key}-{slug}"
        self.run("fetch", settings.git_remote, settings.git_base_branch)
        self.run("checkout", settings.git_base_branch)
        self.run("pull", "--ff-only", settings.git_remote, settings.git_base_branch)
        self.run("checkout", "-b", branch)
        return branch

    def status(self):
        return self.run("status", "--short")

    def diff_stat(self):
        return self.run("diff", "--stat")

    def push(self, jira_key, approved_branch):
        branch = self.current_branch()
        if branch != approved_branch or jira_key not in branch:
            raise RuntimeError("Branch approval mismatch")
        return self.run("push", "-u", settings.git_remote, branch)
