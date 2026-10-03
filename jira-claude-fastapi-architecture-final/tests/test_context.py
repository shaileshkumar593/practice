from app.retrieval.context import build_context

def test_context():
    jira = {
        "key": "PROJ-123",
        "summary": "Create user",
        "status": "Open",
        "description": "Create endpoint",
        "acceptance_criteria": "Return 201",
        "labels": [],
        "linked_issues": [],
    }
    assert "PROJ-123" in build_context(jira, [])
