from fastapi import Header, HTTPException
from app.config import settings

def actor_from_headers(x_user: str | None, authorization: str | None) -> str:
    if settings.oidc_enabled:
        # Production hook: validate JWT signature against the configured JWKS,
        # issuer and audience. Fail closed until that integration is configured.
        raise HTTPException(501, "OIDC validation adapter must be configured")
    if not authorization or authorization != f"Bearer {settings.api_key}":
        raise HTTPException(401, "Unauthorized")
    return x_user or "local-developer"

def require_auth(x_user: str | None = Header(default=None),
                 authorization: str | None = Header(default=None)) -> str:
    return actor_from_headers(x_user, authorization)
