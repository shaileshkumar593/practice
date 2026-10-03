from app.mcp.jira_mcp import JiraMCP

class JiraTool:
    def __init__(self):
        self.mcp = JiraMCP()

    def get_context(self, jira_key):
        return self.mcp.get_issue_context(jira_key)
