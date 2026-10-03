"""
Actual MCP server exposing Jira read tools.

Run independently when Claude Code is configured to connect to this MCP server.
The orchestrator service may also use the same domain adapter internally.
"""

from mcp.server.fastmcp import FastMCP

from app.jira import JiraMCPAdapter

mcp = FastMCP("jira-production")
jira = JiraMCPAdapter()


@mcp.tool()
async def get_jira_issue(jira_key: str) -> dict:
    """Read a Jira issue and return normalized, non-secret context."""
    issue = await jira.get_issue(jira_key)
    return {
        "key": issue.key,
        "summary": issue.summary,
        "description": issue.description,
        "acceptance_criteria": issue.acceptance_criteria,
        "comments": issue.comments,
        "links": issue.links,
    }


if __name__ == "__main__":
    mcp.run()
