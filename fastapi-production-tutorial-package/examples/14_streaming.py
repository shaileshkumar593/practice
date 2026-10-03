from fastapi import FastAPI
from fastapi.responses import StreamingResponse
import asyncio
app=FastAPI()
@app.get("/stream")
async def stream():
    async def generate():
        for i in range(5): yield f"{i}\n"; await asyncio.sleep(.2)
    return StreamingResponse(generate(),media_type="text/plain")
