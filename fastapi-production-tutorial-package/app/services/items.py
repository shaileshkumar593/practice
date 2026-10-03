from sqlalchemy.ext.asyncio import AsyncSession
from app.models.item import Item
from app.repositories.items import ItemRepository
from app.schemas.item import ItemCreate, ItemUpdate
class ItemService:
    def __init__(self, session: AsyncSession): self.repo = ItemRepository(session)
    async def create(self, owner_id: int, data: ItemCreate) -> Item:
        return await self.repo.create(Item(owner_id=owner_id, name=data.name, description=data.description, price=data.price))
    async def list(self, owner_id: int, skip: int, limit: int) -> list[Item]: return await self.repo.list_for_owner(owner_id, skip, limit)
    async def update(self, owner_id: int, item_id: int, data: ItemUpdate) -> Item:
        item = await self.repo.get(item_id)
        if not item or item.owner_id != owner_id: raise LookupError("ITEM_NOT_FOUND")
        for key, value in data.model_dump(exclude_unset=True).items(): setattr(item, key, value)
        await self.repo.session.commit(); await self.repo.session.refresh(item); return item
    async def delete(self, owner_id: int, item_id: int) -> None:
        item = await self.repo.get(item_id)
        if not item or item.owner_id != owner_id: raise LookupError("ITEM_NOT_FOUND")
        await self.repo.delete(item)
