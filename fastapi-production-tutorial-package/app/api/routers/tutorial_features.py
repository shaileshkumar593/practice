from typing import Annotated
import asyncio,json
from datetime import datetime,timezone
from fastapi import APIRouter,Cookie,File,Form,Header,HTTPException,Query,UploadFile
from fastapi.responses import HTMLResponse,JSONResponse,StreamingResponse
from sse_starlette.sse import EventSourceResponse

router=APIRouter(prefix="/tutorial",tags=["tutorial-features"])


@router.get("/parameters/{item_id}")
async def parameters(item_id:int,q:Annotated[str|None,Query(min_length=2,max_length=100)]=None,x_token:Annotated[str|None,Header()]=None,session_id:Annotated[str|None,Cookie()]=None): return {"item_id":item_id,"q":q,"x_token":x_token,"session_id":session_id}
@router.post("/multiple-body")
async def multiple_body(item:dict,metadata:dict|None=None): return {"item":item,"metadata":metadata}
@router.post("/form")
async def form_example(username:Annotated[str,Form()],password:Annotated[str,Form()]): return {"username":username,"password_length":len(password)}
@router.post("/upload")
async def upload(file:UploadFile=File(...)):
    data=await file.read(); return {"filename":file.filename,"content_type":file.content_type,"bytes":len(data)}
@router.get("/errors")
async def errors(): raise HTTPException(status_code=409,detail="Example conflict")
@router.get("/json")
async def json_compatible(): return JSONResponse({"timestamp":datetime.now(timezone.utc).isoformat()})
@router.get("/stream")
async def stream_json_lines():
    async def generator():
        for number in range(5): yield json.dumps({"number":number})+"\n"; await asyncio.sleep(.2)
    return StreamingResponse(generator(),media_type="application/x-ndjson")
@router.get("/events")
async def sse():
    async def event_generator():
        for number in range(5): yield {"event":"tick","data":json.dumps({"number":number})}; await asyncio.sleep(1)
    return EventSourceResponse(event_generator())
@router.get("/html",response_class=HTMLResponse)
async def html_response(): return "<html><body><h1>FastAPI HTML response</h1></body></html>"
