from typing import Annotated
from fastapi import Depends, HTTPException
from fastapi.security import OAuth2PasswordBearer
from jwt import InvalidTokenError
from app.core.security import decode_access_token
from app.db.session import get_db
from app.models.user import User
from app.repositories.users import UserRepository
oauth2_scheme = OAuth2PasswordBearer(tokenUrl="/api/v1/auth/token")
async def get_current_user(token: Annotated[str, Depends(oauth2_scheme)], session=Depends(get_db)) -> User:
    try: user_id = int(decode_access_token(token))
    except (ValueError, InvalidTokenError, TypeError):
        raise HTTPException(status_code=401, detail="Invalid or expired access token", headers={"WWW-Authenticate":"Bearer"})
    user = await UserRepository(session).get_by_id(user_id)
    if not user: raise HTTPException(status_code=401, detail="User not found")
    return user
