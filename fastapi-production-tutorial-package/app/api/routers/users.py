from fastapi import APIRouter, Depends, HTTPException, status
from app.dependencies.auth import get_current_user
from app.models.user import User
from app.schemas.user import UserCreate, UserRead
from app.services.users import UserService
from app.db.session import get_db

router=APIRouter(prefix="/users",tags=["users"])

@router.post("/",response_model=UserRead,status_code=status.HTTP_201_CREATED)
async def create_user(data:UserCreate,session=Depends(get_db)):
    try: return await UserService(session).create(data)
    except ValueError as exc:
        if str(exc)=="EMAIL_ALREADY_EXISTS": raise HTTPException(status_code=409,detail="Email already exists")
        raise

    
@router.get("/me",response_model=UserRead)
async def get_me(current_user:User=Depends(get_current_user)): return current_user
