from pydantic import BaseModel


# -------------------------
# Create Address
# -------------------------
class AddressCreate(BaseModel):
    city: str
    state: str
    country: str


# -------------------------
# Update Address
# -------------------------
class AddressUpdate(BaseModel):
    city: str
    state: str
    country: str


# -------------------------
# Response
# -------------------------
class AddressResponse(BaseModel):
    id: int
    city: str
    state: str
    country: str

    class Config:
        from_attributes = True