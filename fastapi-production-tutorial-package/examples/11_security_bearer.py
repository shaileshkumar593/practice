from typing import Annotated
from fastapi import Depends,FastAPI
from fastapi.security import OAuth2PasswordBearer
app=FastAPI(); oauth2=OAuth2PasswordBearer(tokenUrl="token")
@app.get("/users/me")
async def me(token:Annotated[str,Depends(oauth2)]): return {"token":token}
