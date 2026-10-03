# Security Model

## Threats

| Threat | Control |
|---|---|
| Prompt injection in Jira | Jira content treated as untrusted data |
| Malicious repository file | Agent sandbox + no arbitrary host access |
| Token leakage | Secrets manager + no secret files in context |
| Unauthorized push | RBAC + exact approval + SHA binding |
| Approval replay | TTL + exact commit binding |
| Branch tampering | Branch and SHA validation |
| Personal PAT compromise | GitHub App installation token |
| SSRF | Egress allow-list and proxy policy |
| Arbitrary command execution | Allow-listed test commands |
| Privilege escalation | Separate roles |
| Audit repudiation | Persistent audit events |
| Context poisoning | Trusted-source metadata and retrieval filtering |

## Non-negotiable rules

1. Jira comments are data, not executable instructions.
2. Model output is never treated as a shell command.
3. Git credentials never enter the model context.
4. Production credentials are unavailable to Claude Code.
5. Push and PR creation require separate human approvals.
6. A changed commit invalidates approval.
7. Protected branches remain protected at GitHub.
