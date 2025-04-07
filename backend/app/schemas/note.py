from typing import Optional, List
from pydantic import BaseModel
from datetime import datetime

class NoteBase(BaseModel):
    title: str
    content: str
    type: str
    project_id: Optional[int] = None

class NoteCreate(NoteBase):
    pass

class NoteUpdate(BaseModel):
    title: Optional[str] = None
    content: Optional[str] = None
    type: Optional[str] = None
    project_id: Optional[int] = None

class TagBase(BaseModel):
    name: str

class TagResponse(TagBase):
    id: int

    class Config:
        from_attributes = True

class NoteResponse(NoteBase):
    id: int
    created_at: datetime
    updated_at: datetime
    is_synced: bool
    tags: List[TagResponse] = []

    class Config:
        from_attributes = True 