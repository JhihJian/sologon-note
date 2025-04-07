from typing import Optional, List
from pydantic import BaseModel
from datetime import datetime
from app.schemas.note import NoteResponse

class ProjectBase(BaseModel):
    name: str
    description: Optional[str] = None

class ProjectCreate(ProjectBase):
    pass

class ProjectUpdate(BaseModel):
    name: Optional[str] = None
    description: Optional[str] = None

class ProjectResponse(ProjectBase):
    id: int
    created_at: datetime
    notes: List[NoteResponse] = []

    class Config:
        from_attributes = True 