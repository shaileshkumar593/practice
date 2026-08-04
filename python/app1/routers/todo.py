from fastapi import APIRouter, Depends, HTTPException
from sqlalchemy.orm import Session

import crud
import schemas
from database import get_db

router = APIRouter(
    prefix="/todos",
    tags=["Todos"]
)


@router.get("/", response_model=list[schemas.TodoResponse])
def get_all(db: Session = Depends(get_db)):
    return crud.get_todos(db)


@router.get("/{todo_id}", response_model=schemas.TodoResponse)
def get_one(todo_id: int, db: Session = Depends(get_db)):
    todo = crud.get_todo(db, todo_id)

    if not todo:
        raise HTTPException(404, "Todo not found")

    return todo


@router.post("/", response_model=schemas.TodoResponse)
def create(todo: schemas.TodoCreate,
           db: Session = Depends(get_db)):
    return crud.create_todo(db, todo)


@router.put("/{todo_id}", response_model=schemas.TodoResponse)
def update(todo_id: int,
           todo: schemas.TodoUpdate,
           db: Session = Depends(get_db)):

    result = crud.update_todo(db, todo_id, todo)

    if not result:
        raise HTTPException(404, "Todo not found")

    return result


@router.delete("/{todo_id}")
def delete(todo_id: int,
           db: Session = Depends(get_db)):

    result = crud.delete_todo(db, todo_id)

    if not result:
        raise HTTPException(404, "Todo not found")

    return {"message": "Deleted Successfully"}