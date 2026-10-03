from typing import Annotated
from fastapi import Query
async def pagination(skip: Annotated[int, Query(ge=0)] = 0, limit: Annotated[int, Query(ge=1, le=100)] = 20) -> dict[str,int]:
    return {"skip":skip,"limit":limit}
class TenantContext:
    def __init__(self, tenant_id: str): self.tenant_id = tenant_id
async def get_tenant(x_tenant_id: str | None = None) -> TenantContext:
    return TenantContext(x_tenant_id or "default")
