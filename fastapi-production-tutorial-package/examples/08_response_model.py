from fastapi import FastAPI
from pydantic import BaseModel
app=FastAPI()
class PublicUser(BaseModel): id:int; email:str
@app.get("/users/1",response_model=PublicUser)
async def user(): return {"id":1,"email":"user@example.com","password":"never returned"}
