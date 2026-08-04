from fastapi import FastAPI

from app.database import Base
from app.database import engine

# Register Models
from app.models.student import Student
from app.models.address import Address

from app.routers.student_router import router as student_router
from app.routers.address_router import router as address_router

Base.metadata.create_all(bind=engine)

app = FastAPI(
    title="FastAPI One-To-Many Example",
    version="1.0.0"
)

app.include_router(student_router)
app.include_router(address_router)


@app.get("/")
def home():
    return {
        "message": "FastAPI One-To-Many Relationship API"
    }