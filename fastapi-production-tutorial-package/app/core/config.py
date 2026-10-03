from functools import lru_cache
from pydantic_settings import BaseSettings, SettingsConfigDict

class Settings(BaseSettings):
    app_name: str = "fastapi-production-tutorial"
    environment: str = "local"
    debug: bool = False
    database_url: str = "postgresql+asyncpg://app:app@localhost:5432/app"
    jwt_secret_key: str = "change-me"
    jwt_algorithm: str = "HS256"
    access_token_expire_minutes: int = 30
    cors_origins: str = "http://localhost:3000,http://localhost:8000"
    max_upload_bytes: int = 10 * 1024 * 1024
    model_config = SettingsConfigDict(env_file=".env", extra="ignore")
    @property
    def cors_origin_list(self) -> list[str]:
        return [x.strip() for x in self.cors_origins.split(",") if x.strip()]

@lru_cache
def get_settings() -> Settings:
    return Settings()
