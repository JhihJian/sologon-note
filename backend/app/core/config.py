from pydantic_settings import BaseSettings
from typing import Optional

class Settings(BaseSettings):
    PROJECT_NAME: str = "Sologon"
    VERSION: str = "0.1.0"
    API_V1_STR: str = "/api/v1"
    
    # 数据库配置
    SQLALCHEMY_DATABASE_URL: str = "sqlite:///./sologon.db"
    
    # JWT配置
    SECRET_KEY: str = "your-secret-key-here"  # 在生产环境中应该使用环境变量
    ALGORITHM: str = "HS256"
    ACCESS_TOKEN_EXPIRE_MINUTES: int = 30
    
    # GitHub配置
    GITHUB_TOKEN: Optional[str] = None
    GITHUB_REPO: Optional[str] = None
    
    class Config:
        case_sensitive = True

settings = Settings() 