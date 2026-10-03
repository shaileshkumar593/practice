from pydantic import BaseModel, ConfigDict, Field
class ItemCreate(BaseModel):
    name: str = Field(min_length=1, max_length=200)
    description: str | None = Field(default=None, max_length=5000)
    price: float = Field(gt=0)
class ItemRead(BaseModel):
    model_config = ConfigDict(from_attributes=True)
    id: int
    owner_id: int
    name: str
    description: str | None
    price: float
class ItemUpdate(BaseModel):
    name: str | None = Field(default=None, min_length=1, max_length=200)
    description: str | None = Field(default=None, max_length=5000)
    price: float | None = Field(default=None, gt=0)
