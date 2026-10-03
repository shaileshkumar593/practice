from typing import Annotated
from fastapi import FastAPI,File,Form,UploadFile
app=FastAPI()
@app.post("/upload")
async def upload(username:Annotated[str,Form()],file:UploadFile=File(...)): return {"username":username,"filename":file.filename}
