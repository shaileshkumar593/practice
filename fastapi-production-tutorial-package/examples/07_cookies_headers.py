from typing import Annotated
from fastapi import Cookie,FastAPI,Header
app=FastAPI()
@app.get("/context")
async def context(user_agent:Annotated[str|None,Header()]=None,session:Annotated[str|None,Cookie()]=None): return {"user_agent":user_agent,"session":session}
