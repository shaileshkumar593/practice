from fastapi import FastAPI, Depends
from app.dependencies import require_api_key
from app.routes.work_items import router

app = FastAPI(
    title="Jira Claude Coding Agent",
    version="2.0.0",
    description="Jira MCP -> Hybrid RAG -> Claude Code -> Human Gate -> GitHub PR",
    dependencies=[Depends(require_api_key)],
)

app.include_router(router, prefix="/api/v1")

@app.get("/health")
def health():
    return {"status": "ok"}
