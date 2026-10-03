from fastapi import FastAPI
from sse_starlette.sse import EventSourceResponse
import asyncio
app=FastAPI()
@app.get("/events")
async def events():
    async def generate():
        for i in range(5): yield {"event":"tick","data":str(i)}; await asyncio.sleep(1)
    return EventSourceResponse(generate())
