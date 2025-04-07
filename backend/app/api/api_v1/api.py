from fastapi import APIRouter
from app.api.api_v1.endpoints import notes, projects, tags

api_router = APIRouter()
api_router.include_router(notes.router, prefix="/notes", tags=["notes"])
api_router.include_router(projects.router, prefix="/projects", tags=["projects"])
api_router.include_router(tags.router, prefix="/tags", tags=["tags"]) 