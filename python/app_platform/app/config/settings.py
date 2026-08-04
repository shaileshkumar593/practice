from functools import lru_cache

from pydantic_settings import BaseSettings
from pydantic import ConfigDict


class Settings(BaseSettings):

    APP_NAME: str = "Enterprise AI Platform"

    APP_VERSION: str = "1.0.0"

    DEBUG: bool = True

    HOST: str = "0.0.0.0"

    PORT: int = 8000

    API_PREFIX: str = "/api/v1"

    SECRET_KEY: str = "change-this-in-production"

    ACCESS_TOKEN_EXPIRE_MINUTES: int = 60

    DATABASE_URL: str = (
        "postgresql+psycopg://postgres:postgres@localhost:5432/enterprise_ai"
    )

    REDIS_URL: str = (
        "redis://localhost:6379/0"
    )

    KAFKA_BOOTSTRAP_SERVERS: str = (
        "localhost:9092"
    )

    OPENAI_API_KEY: str = ""

    ELASTICSEARCH_URL: str = (
        "http://localhost:9200"
    )

    VECTOR_DB_URL: str = (
        "http://localhost:6333"
    )

    LOG_LEVEL: str = "INFO"

    model_config = ConfigDict(
        env_file=".env",
        extra="ignore"
    )


@lru_cache
def get_settings():

    return Settings()


settings = get_settings()