from fastapi import FastAPI

from app.database import Base, engine

# Import models so SQLAlchemy registers them
from app.models.user import User
from app.models.profile import Profile

from app.routers.user_router import router as user_router
from app.routers.profile_router import router as profile_router

Base.metadata.create_all(bind=engine)

app = FastAPI(
    title="FastAPI One-To-One Example",
    version="1.0.0"
)

app.include_router(user_router)
app.include_router(profile_router)


@app.get("/")
def home():
    return {
        "message": "FastAPI One-To-One Relationship API"
    }