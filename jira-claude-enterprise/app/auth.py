from dataclasses import dataclass
from functools import lru_cache
from typing import Any

import httpx
import jwt
from fastapi import Header, HTTPException

from app.config import settings


@dataclass(frozen=True)
class Principal:
    subject: str
    roles: frozenset[str]
    scopes: frozenset[str]


ROLE_PERMISSIONS = {
    "developer": {
        "workitem:start",
        "workitem:inspect",
        "workitem:ready",
    },
    "reviewer": {
        "workitem:inspect",
        "workitem:approve_push",
        "workitem:approve_pr",
    },
    "release-manager": {
        "workitem:inspect",
        "workitem:push",
    },
    "admin": {"*"},
}


def _permissions(principal: Principal) -> set[str]:
    permissions: set[str] = set()
    for role in principal.roles:
        permissions.update(ROLE_PERMISSIONS.get(role, set()))
    permissions.update(principal.scopes)
    return permissions


def require_permission(principal: Principal, permission: str) -> None:
    if "*" not in _permissions(principal) and permission not in _permissions(principal):
        raise HTTPException(status_code=403, detail="Insufficient permissions")


@lru_cache(maxsize=1)
def _jwks() -> dict[str, Any]:
    if not settings.oidc_jwks_url:
        raise RuntimeError("OIDC_JWKS_URL is required when OIDC is enabled")
    with httpx.Client(timeout=10) as client:
        response = client.get(settings.oidc_jwks_url)
        response.raise_for_status()
        return response.json()


def _oidc_principal(token: str) -> Principal:
    header = jwt.get_unverified_header(token)
    key = next(
        (
            item
            for item in _jwks().get("keys", [])
            if item.get("kid") == header.get("kid")
        ),
        None,
    )
    if not key:
        raise HTTPException(status_code=401, detail="Unknown signing key")

    try:
        payload = jwt.decode(
            token,
            key,
            algorithms=["RS256"],
            issuer=settings.oidc_issuer,
            audience=settings.oidc_audience,
        )
    except jwt.PyJWTError as exc:
        raise HTTPException(status_code=401, detail="Invalid identity token") from exc

    roles = payload.get("roles", payload.get("groups", []))
    if isinstance(roles, str):
        roles = [roles]

    scopes = payload.get("scope", "")
    scopes = scopes.split() if isinstance(scopes, str) else scopes

    return Principal(
        subject=str(payload.get("sub", "unknown")),
        roles=frozenset(roles),
        scopes=frozenset(scopes),
    )


def current_principal(
    authorization: str | None = Header(default=None),
    x_user: str | None = Header(default=None),
) -> Principal:
    if settings.oidc_enabled:
        if not authorization or not authorization.startswith("Bearer "):
            raise HTTPException(status_code=401, detail="Bearer token required")
        return _oidc_principal(authorization.removeprefix("Bearer ").strip())

    if authorization != f"Bearer {settings.api_key}":
        raise HTTPException(status_code=401, detail="Unauthorized")

    return Principal(
        subject=x_user or "local-developer",
        roles=frozenset({"admin"}),
        scopes=frozenset({"*"}),
    )
