from fastapi import FastAPI,WebSocket
app=FastAPI()
@app.websocket("/ws")
async def websocket(ws:WebSocket):
    await ws.accept()
    while True: await ws.send_text(await ws.receive_text())
