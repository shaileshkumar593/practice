from fastapi import APIRouter, Depends
from sqlalchemy.orm import Session

from app.database import get_db

from app.schemas.student import (
    StudentCreate,
    StudentUpdate,
    StudentResponse
)

from app.services.student_service import StudentService

router = APIRouter(
    prefix="/students",
    tags=["Students"]
)


@router.post(
    "/",
    response_model=StudentResponse,
    status_code=201
)
def create_student(
    student: StudentCreate,
    db: Session = Depends(get_db)
):
    return StudentService.create_student(
        db,
        student
    )


@router.get(
    "/",
    response_model=list[StudentResponse]
)
def get_students(
    db: Session = Depends(get_db)
):
    return StudentService.get_students(db)


@router.get(
    "/{student_id}",
    response_model=StudentResponse
)
def get_student(
    student_id: int,
    db: Session = Depends(get_db)
):
    return StudentService.get_student(
        db,
        student_id
    )


@router.put(
    "/{student_id}",
    response_model=StudentResponse
)
def update_student(
    student_id: int,
    student: StudentUpdate,
    db: Session = Depends(get_db)
):
    return StudentService.update_student(
        db,
        student_id,
        student
    )


@router.delete("/{student_id}")
def delete_student(
    student_id: int,
    db: Session = Depends(get_db)
):
    return StudentService.delete_student(
        db,
        student_id
    )