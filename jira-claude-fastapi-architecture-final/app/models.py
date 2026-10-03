from pydantic import BaseModel

class ReadyRequest(BaseModel):
    test_command: str | None = None
    test_output: str | None = None
    notes: str | None = None

class ApprovalRequest(BaseModel):
    phrase: str

class PRRequest(BaseModel):
    title: str | None = None
    body: str | None = None
