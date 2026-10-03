import time
from pathlib import Path

import httpx
import jwt

from app.config import settings


class GitHubAppError(RuntimeError):
    pass


class GitHubAppClient:
    """
    GitHub App authentication.

    Uses a short-lived installation token instead of a developer PAT.
    """

    API = "https://api.github.com"

    def _app_jwt(self) -> str:
        if not settings.github_app_id or not settings.github_private_key_path:
            raise GitHubAppError("GitHub App credentials are not configured")

        key = Path(settings.github_private_key_path).read_text(encoding="utf-8")
        now = int(time.time())
        return jwt.encode(
            {
                "iat": now - 30,
                "exp": now + 540,
                "iss": str(settings.github_app_id),
            },
            key,
            algorithm="RS256",
        )

    def installation_token(self) -> str:
        if not settings.github_installation_id:
            raise GitHubAppError("GITHUB_INSTALLATION_ID is required")

        token = self._app_jwt()
        url = (
            f"{self.API}/app/installations/"
            f"{settings.github_installation_id}/access_tokens"
        )
        with httpx.Client(timeout=15) as client:
            response = client.post(
                url,
                headers={
                    "Authorization": f"Bearer {token}",
                    "Accept": "application/vnd.github+json",
                },
            )
            response.raise_for_status()
            return response.json()["token"]

    def create_pull_request(
        self,
        head: str,
        base: str,
        title: str,
        body: str,
    ) -> dict:
        token = self.installation_token()
        owner, repo = settings.github_repository.split("/", 1)
        url = f"{self.API}/repos/{owner}/{repo}/pulls"

        with httpx.Client(timeout=20) as client:
            response = client.post(
                url,
                headers={
                    "Authorization": f"Bearer {token}",
                    "Accept": "application/vnd.github+json",
                },
                json={
                    "title": title,
                    "head": head,
                    "base": base,
                    "body": body,
                },
            )
            response.raise_for_status()
            return response.json()
