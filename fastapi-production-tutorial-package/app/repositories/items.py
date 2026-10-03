from sqlalchemy import select
from sqlalchemy.ext.asyncio import AsyncSession
from app.models.item import Item
class ItemRepository:
    def __init__(self, session: AsyncSession): self.session = session
    async def get(self, item_id: int) -> Item | None:
        return (await self.session.execute(select(Item).where(Item.id == item_id))).scalar_one_or_none()
    async def list_for_owner(self, owner_id: int, skip: int, limit: int) -> list[Item]:
        return list((await self.session.execute(select(Item).where(Item.owner_id == owner_id).offset(skip).limit(limit))).scalars().all())
    async def create(self, item: Item) -> Item:
        self.session.add(item); await self.session.commit(); await self.session.refresh(item); return item
    async def delete(self, item: Item) -> None:
        await self.session.delete(item); await self.session.commit()
