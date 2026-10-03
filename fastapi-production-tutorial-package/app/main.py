from contextlib import asynccontextmanager
from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from fastapi.staticfiles import StaticFiles
from app.api.routers import auth,items,streaming,tutorial_features,users,websocket
from app.core.config import get_settings
from app.middleware.logging import AccessLogMiddleware
from app.middleware.request_id import RequestIDMiddleware
settings=get_settings()
@asynccontextmanager
async def lifespan(app:FastAPI): yield
app=FastAPI(title=settings.app_name,version="1.0.0",description="Production-oriented FastAPI tutorial reference application.",docs_url="/docs",redoc_url="/redoc",openapi_url="/openapi.json",lifespan=lifespan)
app.add_middleware(RequestIDMiddleware); app.add_middleware(AccessLogMiddleware)
app.add_middleware(CORSMiddleware,allow_origins=settings.cors_origin_list,allow_credentials=True,allow_methods=["*"],allow_headers=["*"])
app.mount("/static",StaticFiles(directory="app/static"),name="static")
app.include_router(auth.router,prefix="/api/v1"); app.include_router(users.router,prefix="/api/v1"); app.include_router(items.router,prefix="/api/v1"); app.include_router(tutorial_features.router); app.include_router(streaming.router); app.include_router(websocket.router)
@app.get("/health",tags=["operations"])
async def health(): return {"status":"ok"}
@app.get("/ready",tags=["operations"])
async def ready(): return {"status":"ready"}
