from fastapi import APIRouter, WebSocket, WebSocketDisconnect

router=APIRouter(tags=["websocket"])

@router.websocket("/ws")
async def websocket_endpoint(websocket:WebSocket):
    await websocket.accept()
    try:
        while True: await websocket.send_text(f"echo: {await websocket.receive_text()}")
    except WebSocketDisconnect: pass
