from sqlalchemy.ext.asyncio import AsyncSession
from app.core.security import hash_password
from app.models.user import User
from app.repositories.users import UserRepository
from app.schemas.user import UserCreate
class UserService:
    def __init__(self, session: AsyncSession): self.repo = UserRepository(session)
    async def create(self, data: UserCreate) -> User:
        if await self.repo.get_by_email(data.email): raise ValueError("EMAIL_ALREADY_EXISTS")
        return await self.repo.create(User(email=data.email, full_name=data.full_name, hashed_password=hash_password(data.password)))
