import requests
from app.config import settings

class JiraMCP:
    # MCP-shaped boundary. Replace internals with your real Jira MCP server/connector
    # when available; Claude should consume Jira through tools, not business logic.

    def get_issue(self, key: str):
        url = f"{settings.jira_base_url.rstrip('/')}/rest/api/3/issue/{key}"
        r = requests.get(
            url,
            auth=(settings.jira_email, settings.jira_api_token),
            headers={"Accept": "application/json"},
            params={"expand": "comments,issuelinks"},
            timeout=30,
        )
        r.raise_for_status()
        fields = r.json()["fields"]

        comments = [
            adf_to_text(x.get("body"))
            for x in fields.get("comment", {}).get("comments", [])
        ]

        links = []
        for link in fields.get("issuelinks", []):
            for side in ("inwardIssue", "outwardIssue"):
                issue = link.get(side)
                if issue:
                    links.append({
                        "key": issue.get("key"),
                        "summary": issue.get("fields", {}).get("summary", ""),
                    })

        description = adf_to_text(fields.get("description"))

        marker = description.lower().find("acceptance criteria")
        acceptance = description[marker:] if marker >= 0 else ""

        return {
            "key": key,
            "summary": fields.get("summary", ""),
            "description": description,
            "status": (fields.get("status") or {}).get("name", ""),
            "issue_type": (fields.get("issuetype") or {}).get("name", ""),
            "priority": (fields.get("priority") or {}).get("name", ""),
            "labels": fields.get("labels", []),
            "comments": [x for x in comments if x],
            "linked_issues": links,
            "acceptance_criteria": acceptance,
        }

    def get_comments(self, key):
        return self.get_issue(key)["comments"]

    def get_linked_issues(self, key):
        return self.get_issue(key)["linked_issues"]

    def get_issue_context(self, key):
        return self.get_issue(key)

def adf_to_text(node):
    if not node:
        return ""
    if isinstance(node, str):
        return node
    if isinstance(node, list):
        return "\n".join(x for x in (adf_to_text(v) for v in node) if x)
    if isinstance(node, dict):
        if node.get("type") == "text":
            return node.get("text", "")
        return adf_to_text(node.get("content", []))
    return ""
