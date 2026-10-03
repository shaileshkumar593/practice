from fastapi import FastAPI
from pydantic import BaseModel
app=FastAPI()
class Image(BaseModel): url:str
class Item(BaseModel): name:str; images:list[Image]=[]
@app.post("/items")
async def create(item:Item): return item
