# Claude Code Policy

Architecture:
Jira -> Jira MCP -> Hybrid RAG -> Claude Code -> Jira/Git/Test tools -> Human Review -> Gate -> GitHub PR -> CI/CD.

Read `.claude/context/<JIRA-ID>.md` before implementation.

Implement only the Jira requirement. Inspect the repository and tests. Follow existing patterns.

Run formatting, tests, static checks and build checks as appropriate.

After implementation show:
- requirements
- acceptance criteria mapping
- changed files
- architecture decisions
- API/DB/event changes
- tests and output
- git diff
- security
- performance
- risks

Then STOP and say:

WAITING FOR HUMAN REVIEW

Never before explicit approval:
- git push
- create PR
- merge
- deploy
- modify main/master

Only this phrase authorizes push:

APPROVED TO PUSH <JIRA-ID>

The ID must match the current work item.
