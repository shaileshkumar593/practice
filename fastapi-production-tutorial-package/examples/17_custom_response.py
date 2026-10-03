from fastapi import FastAPI
from fastapi.responses import HTMLResponse,JSONResponse
app=FastAPI()
@app.get("/html",response_class=HTMLResponse)
async def html(): return "<h1>Hello</h1>"
@app.get("/json")
async def json(): return JSONResponse({"ok":True})
