import pytest
@pytest.mark.asyncio
async def test_health(client):
    r=await client.get("/health"); assert r.status_code==200; assert r.json()["status"]=="ok"
@pytest.mark.asyncio
async def test_query_header(client):
    r=await client.get("/tutorial/parameters/5",params={"q":"hello"},headers={"X-Token":"abc"})
    assert r.status_code==200 and r.json()["item_id"]==5
@pytest.mark.asyncio
async def test_validation(client):
    r=await client.get("/tutorial/parameters/5",params={"q":"x"}); assert r.status_code==422
