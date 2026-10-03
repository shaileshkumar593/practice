from pydantic_settings import BaseSettings, SettingsConfigDict


class Settings(BaseSettings):
    app_env: str = "development"
    log_level: str = "INFO"

    database_url: str
    api_key: str = "change-me-in-development"
    workspace_root: str = "/workspace/repository"

    oidc_enabled: bool = True
    oidc_issuer: str = ""
    oidc_audience: str = ""
    oidc_jwks_url: str = ""
    rbac_required: bool = True

    jira_base_url: str = ""
    jira_email: str = ""
    jira_api_token: str = ""

    github_app_id: int | None = None
    github_installation_id: int | None = None
    github_private_key_path: str = ""
    github_repository: str = ""
    github_base_branch: str = "main"

    retrieval_backend: str = "hybrid"
    opensearch_url: str = ""
    opensearch_index: str = "repository-code"
    pgvector_enabled: bool = True
    embedding_provider: str = "stub"
    embedding_model: str = "text-embedding-3-small"
    reranker_provider: str = "stub"
    reranker_model: str = ""

    approval_ttl_seconds: int = 1800
    max_retrieval_results: int = 12
    test_timeout_seconds: int = 300

    model_config = SettingsConfigDict(
        env_file=".env",
        env_file_encoding="utf-8",
        extra="ignore",
    )


settings = Settings()
