from fastapi import APIRouter, HTTPException
from app.models import ReadyRequest, ApprovalRequest, PRRequest
from app.services.orchestrator import Orchestrator

router = APIRouter(tags=["work-items"])
service = Orchestrator()

@router.post("/work-items/{jira_key}/start")
def start(jira_key: str):
    try:
        return service.start(jira_key)
    except Exception as e:
        raise HTTPException(status_code=502, detail=str(e))

@router.get("/work-items/{jira_key}")
def inspect(jira_key: str):
    item = service.inspect(jira_key)
    if not item:
        raise HTTPException(status_code=404, detail="Not found")
    return item

@router.post("/work-items/{jira_key}/ready-for-review")
def ready(jira_key: str, request: ReadyRequest):
    try:
        return service.ready(jira_key, request.model_dump())
    except ValueError as e:
        raise HTTPException(status_code=409, detail=str(e))

@router.post("/work-items/{jira_key}/approve")
def approve(jira_key: str, request: ApprovalRequest):
    try:
        return service.approve(jira_key, request.phrase)
    except ValueError as e:
        raise HTTPException(status_code=409, detail=str(e))

@router.post("/work-items/{jira_key}/push")
def push(jira_key: str):
    try:
        return service.push(jira_key)
    except ValueError as e:
        raise HTTPException(status_code=409, detail=str(e))

@router.post("/work-items/{jira_key}/create-pr")
def create_pr(jira_key: str, request: PRRequest):
    try:
        return service.create_pr(jira_key, request.title, request.body)
    except ValueError as e:
        raise HTTPException(status_code=409, detail=str(e))
