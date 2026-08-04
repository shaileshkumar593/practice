from pydantic import BaseModel, EmailStr

from app.schemas.profile import (
    ProfileCreate,
    ProfileResponse,
    ProfileUpdate,
)


# --------------------------------
# Create User
# --------------------------------
class UserCreate(BaseModel):
    name: str
    email: EmailStr
    profile: ProfileCreate


# --------------------------------
# Update User
# --------------------------------
class UserUpdate(BaseModel):
    name: str
    email: EmailStr
    profile: ProfileUpdate


# --------------------------------
# Response
# --------------------------------
class UserResponse(BaseModel):
    id: int
    name: str
    email: EmailStr
    profile: ProfileResponse

    class Config:
        from_attributes = True