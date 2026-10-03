from dataclasses import dataclass
import httpx
from app.config import settings

@dataclass(frozen=True)
class JiraIssue:
    key: str
    summary: str
    description: str
    acceptance_criteria: str
    comments: list[str]
    links: list[str]

class JiraMCPAdapter:
    """Strict Jira boundary. In production this is the adapter behind a real MCP server."""

    async def get_issue(self, key: str) -> JiraIssue:
        if not settings.jira_base_url:
            return JiraIssue(key, "Demo issue", "Demo description", "Add tests", [], [])
        url = f"{settings.jira_base_url.rstrip('/')}/rest/api/3/issue/{key}"
        auth = (settings.jira_email, settings.jira_api_token)
        async with httpx.AsyncClient(timeout=15) as client:
            r = await client.get(url, auth=auth, params={"fields": "summary,description,comment,issuelinks"})
            r.raise_for_status()
            data = r.json()
        fields = data.get("fields", {})
        comments = [c.get("body", "") for c in fields.get("comment", {}).get("comments", [])]
        links = []
        for item in fields.get("issuelinks", []):
            outward = item.get("outwardIssue") or item.get("inwardIssue")
            if outward and outward.get("key"):
                links.append(outward["key"])
        return JiraIssue(key, fields.get("summary", ""), str(fields.get("description", "")),
                         "", comments, links)
