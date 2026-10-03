from pathlib import Path
import json
from threading import Lock
from datetime import datetime, timezone

ROOT = Path(".agent_data")
ROOT.mkdir(exist_ok=True)
FILE = ROOT / "work_items.json"
LOCK = Lock()

def now():
    return datetime.now(timezone.utc).isoformat()

def load():
    if not FILE.exists():
        return {}
    return json.loads(FILE.read_text())

def get(key):
    with LOCK:
        return load().get(key)

def save(item):
    with LOCK:
        data = load()
        data[item["jira_key"]] = item
        FILE.write_text(json.dumps(data, indent=2))
        return item
