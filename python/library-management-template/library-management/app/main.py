from fastapi import FastAPI
from app.api.books import router as books_router
from app.api.members import router as members_router
from app.api.borrow import router as borrow_router

app = FastAPI(title="Library Management")

app.include_router(books_router)
app.include_router(members_router)
app.include_router(borrow_router)
