from pydantic import BaseModel
from pydantic import EmailStr

from typing import List

from app.schemas.address import (
    AddressCreate,
    AddressResponse,
    AddressUpdate
)


# -------------------------
# Create Student
# -------------------------
class StudentCreate(BaseModel):

    name: str

    email: EmailStr

    addresses: List[AddressCreate]


# -------------------------
# Update Student
# -------------------------
class StudentUpdate(BaseModel):

    name: str

    email: EmailStr

    addresses: List[AddressUpdate]


# -------------------------
# Response
# -------------------------
class StudentResponse(BaseModel):

    id: int

    name: str

    email: EmailStr

    addresses: List[AddressResponse]

    class Config:
        from_attributes = True