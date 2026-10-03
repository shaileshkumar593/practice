from typing import Annotated
from fastapi import APIRouter, BackgroundTasks, Depends, HTTPException, Path, Query, Response, status
from app.db.session import get_db
from app.dependencies.auth import get_current_user
from app.models.user import User
from app.schemas.item import ItemCreate, ItemRead, ItemUpdate
from app.services.items import ItemService
from app.workers.tasks import write_audit_log
router=APIRouter(prefix="/items",tags=["items"])
@router.post("/",response_model=ItemRead,status_code=status.HTTP_201_CREATED)
async def create_item(data:ItemCreate,background_tasks:BackgroundTasks,current_user:User=Depends(get_current_user),session=Depends(get_db)):
    item=await ItemService(session).create(current_user.id,data); background_tasks.add_task(write_audit_log,"item.created",item.id); return item
@router.get("/",response_model=list[ItemRead])
async def list_items(skip:Annotated[int,Query(ge=0)]=0,limit:Annotated[int,Query(ge=1,le=100)]=20,current_user:User=Depends(get_current_user),session=Depends(get_db)):
    return await ItemService(session).list(current_user.id,skip,limit)
@router.patch("/{item_id}",response_model=ItemRead)
async def update_item(data:ItemUpdate,item_id:Annotated[int,Path(gt=0)],current_user:User=Depends(get_current_user),session=Depends(get_db)):
    try:return await ItemService(session).update(current_user.id,item_id,data)
    except LookupError:raise HTTPException(status_code=404,detail="Item not found")
@router.delete("/{item_id}",status_code=204)
async def delete_item(item_id:Annotated[int,Path(gt=0)],current_user:User=Depends(get_current_user),session=Depends(get_db)):
    try:await ItemService(session).delete(current_user.id,item_id)
    except LookupError:raise HTTPException(status_code=404,detail="Item not found")
    return Response(status_code=204)
