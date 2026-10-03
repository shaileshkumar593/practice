from typing import Annotated
from fastapi import FastAPI,Query
app=FastAPI()
@app.get("/search")
async def search(q:Annotated[str,Query(min_length=3,max_length=50)]): return {"q":q}
