from fastapi import FastAPI

import models
from database import Base, engine
from routers.todo import router as todo_router

Base.metadata.create_all(bind=engine)

app = FastAPI(
    title="Todo API"
)

app.include_router(todo_router)


@app.get("/")
def home():
    return {
        "message": "FastAPI CRUD is running"
    }