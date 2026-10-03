from typing import Annotated
from fastapi import Depends,FastAPI
app=FastAPI()
async def common_query(q:str|None=None): return q
@app.get("/items")
async def items(q:Annotated[str|None,Depends(common_query)]): return {"q":q}
