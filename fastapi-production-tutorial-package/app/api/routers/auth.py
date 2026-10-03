from typing import Annotated
from fastapi import APIRouter, Depends, HTTPException
from fastapi.security import OAuth2PasswordRequestForm
from app.core.security import create_access_token, verify_password
from app.db.session import get_db
from app.repositories.users import UserRepository
from app.schemas.auth import Token
router=APIRouter(prefix="/auth",tags=["authentication"])
@router.post("/token",response_model=Token)
async def login(form: Annotated[OAuth2PasswordRequestForm,Depends()],session=Depends(get_db)):
    user=await UserRepository(session).get_by_email(form.username)
    if not user or not verify_password(form.password,user.hashed_password): raise HTTPException(status_code=401,detail="Incorrect username or password",headers={"WWW-Authenticate":"Bearer"})
    return Token(access_token=create_access_token(str(user.id)))
