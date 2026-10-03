from datetime import datetime, timedelta, timezone
import jwt
from pwdlib import PasswordHash
from app.core.config import get_settings
password_hash = PasswordHash.recommended()

def hash_password(password: str) -> str:
    return password_hash.hash(password)

def verify_password(password: str, hashed: str) -> bool:
    return password_hash.verify(password, hashed)

def create_access_token(subject: str) -> str:
    s = get_settings()
    expires = datetime.now(timezone.utc) + timedelta(minutes=s.access_token_expire_minutes)
    return jwt.encode({"sub": subject, "exp": expires}, s.jwt_secret_key, algorithm=s.jwt_algorithm)

def decode_access_token(token: str) -> str:
    s = get_settings()
    payload = jwt.decode(token, s.jwt_secret_key, algorithms=[s.jwt_algorithm])
    subject = payload.get("sub")
    if not subject:
        raise ValueError("Token subject missing")
    return str(subject)
