from pydantic_settings import BaseSettings, SettingsConfigDict

class Settings(BaseSettings):
    api_key: str = "change-me"
    jira_base_url: str
    jira_email: str
    jira_api_token: str
    workspace_root: str
    git_remote: str = "origin"
    git_base_branch: str = "main"
    github_repo: str | None = None
    lexical_weight: float = 0.35
    embedding_weight: float = 0.35
    symbol_weight: float = 0.20
    metadata_weight: float = 0.10
    top_k: int = 20

    model_config = SettingsConfigDict(env_file=".env", extra="ignore")

settings = Settings()
