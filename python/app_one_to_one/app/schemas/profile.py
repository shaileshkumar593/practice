from pydantic import BaseModel


# -----------------------------
# Create Profile
# -----------------------------
class ProfileCreate(BaseModel):
    age: int
    city: str
    phone: str


# -----------------------------
# Update Profile
# -----------------------------
class ProfileUpdate(BaseModel):
    age: int
    city: str
    phone: str


# -----------------------------
# Response
# -----------------------------
class ProfileResponse(BaseModel):
    id: int
    age: int
    city: str
    phone: str

    class Config:
        from_attributes = True