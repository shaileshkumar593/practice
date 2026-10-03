def build_context(jira, results):
    lines = [
        f"# Claude Context Pack — {jira['key']}",
        "",
        "## Jira",
        f"Summary: {jira['summary']}",
        f"Status: {jira['status']}",
        f"Priority: {jira['priority']}",
        "",
        "### Description",
        jira["description"],
        "",
        "### Acceptance Criteria",
        jira["acceptance_criteria"],
        "",
        "### Linked Issues",
    ]

    for x in jira["linked_issues"]:
        lines.append(f"- {x['key']}: {x['summary']}")

    lines += ["", "## Retrieved Repository Context"]

    for i, r in enumerate(results, 1):
        c = r.chunk
        lines += [
            "",
            f"### {i}. {c.path}",
            f"Language: {c.language}",
            f"Symbol: {c.symbol or 'chunk'}",
            f"Lines: {c.start_line}-{c.end_line}",
            f"Final score: {r.final_score:.4f}",
            f"Reasons: {', '.join(r.reasons)}",
            "```",
            c.text,
            "```",
        ]

    return "\n".join(lines)
