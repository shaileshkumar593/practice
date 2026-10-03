from pathlib import Path
from app.config import settings
from app.store import get, save, now
from app.tools.jira_tool import JiraTool
from app.tools.git_tool import GitTool
from app.retrieval.indexer import RepositoryIndexer
from app.retrieval.hybrid import HybridRetriever
from app.retrieval.context import build_context

class Orchestrator:
    def __init__(self):
        self.jira = JiraTool()
        self.git = GitTool()

    def start(self, jira_key):
        if get(jira_key):
            return get(jira_key)

        jira = self.jira.get_context(jira_key)
        branch = self.git.create_branch(jira_key, jira["summary"])

        chunks = RepositoryIndexer().index(Path(settings.workspace_root))
        retriever = HybridRetriever(chunks)

        query = "\n".join([
            jira["summary"],
            jira["description"],
            jira["acceptance_criteria"],
            " ".join(jira["labels"]),
            " ".join(x["summary"] for x in jira["linked_issues"]),
        ])

        results = retriever.search(query, settings.top_k)
        context = build_context(jira, results)

        context_file = Path(settings.workspace_root) / ".claude" / "context" / f"{jira_key}.md"
        context_file.parent.mkdir(parents=True, exist_ok=True)
        context_file.write_text(context, encoding="utf-8")

        return save({
            "jira_key": jira_key,
            "summary": jira["summary"],
            "branch": branch,
            "status": "IMPLEMENTING",
            "context_file": str(context_file),
            "approval": False,
            "approval_branch": None,
            "validation": {},
            "retrieval": {
                "indexed_chunks": len(chunks),
                "returned_chunks": len(results),
                "results": [
                    {
                        "path": r.chunk.path,
                        "symbol": r.chunk.symbol,
                        "language": r.chunk.language,
                        "score": r.final_score,
                        "reasons": r.reasons,
                    }
                    for r in results
                ],
            },
            "created_at": now(),
            "updated_at": now(),
        })

    def ready(self, jira_key, validation):
        item = get(jira_key)
        if not item:
            raise ValueError("Work item not found")
        item["status"] = "READY_FOR_REVIEW"
        item["validation"] = validation
        item["updated_at"] = now()
        return save(item)

    def approve(self, jira_key, phrase):
        expected = f"APPROVED TO PUSH {jira_key}"
        if phrase.strip() != expected:
            raise ValueError(f"Expected exactly: {expected}")

        item = get(jira_key)
        if not item or item["status"] != "READY_FOR_REVIEW":
            raise ValueError("Work item is not ready for approval")

        branch = self.git.current_branch()
        if branch != item["branch"]:
            raise ValueError("Current branch does not match work item")

        item["approval"] = True
        item["approval_branch"] = branch
        item["status"] = "APPROVED"
        item["updated_at"] = now()
        return save(item)

    def push(self, jira_key):
        item = get(jira_key)
        if not item or not item["approval"]:
            raise ValueError("Human approval is required")

        output = self.git.push(jira_key, item["approval_branch"])
        item["status"] = "PUSHED"
        item["push_output"] = output
        item["updated_at"] = now()
        return save(item)

    def create_pr(self, jira_key, title=None, body=None):
        item = get(jira_key)
        if not item or item["status"] != "PUSHED":
            raise ValueError("Push must complete before PR creation")

        title = title or f"{jira_key}: {item['summary']}"
        body = body or f"Implements {jira_key}."

        import subprocess
        p = subprocess.run(
            ["gh", "pr", "create",
             "--base", settings.git_base_branch,
             "--head", item["branch"],
             "--title", title,
             "--body", body],
            cwd=settings.workspace_root,
            capture_output=True,
            text=True,
        )

        if p.returncode:
            raise ValueError(p.stderr.strip())

        item["status"] = "PR_CREATED"
        item["pr_url"] = p.stdout.strip()
        item["updated_at"] = now()
        return save(item)

    def inspect(self, jira_key):
        item = get(jira_key)
        if not item:
            return None
        result = dict(item)
        result["current_branch"] = self.git.current_branch()
        result["git_status"] = self.git.status()
        result["git_diff_stat"] = self.git.diff_stat()
        return result
