import os
os.environ.setdefault("JIRA_BASE_URL", "https://example.atlassian.net")
os.environ.setdefault("JIRA_EMAIL", "test@example.com")
os.environ.setdefault("JIRA_API_TOKEN", "test")
os.environ.setdefault("WORKSPACE_ROOT", ".")

from fastapi.testclient import TestClient
from app.main import app

def test_health():
    r = TestClient(app).get("/health")
    assert r.status_code == 200
