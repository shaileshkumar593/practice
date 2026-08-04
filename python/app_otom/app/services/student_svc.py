from sqlalchemy.orm import Session
from fastapi import HTTPException

from app.repository.student_repository import StudentRepository
from app.schemas.student import (
    StudentCreate,
    StudentUpdate
)
from app.models.student import Student


class StudentService:

    @staticmethod
    def create_student(
        db: Session,
        student: StudentCreate
    ):

        existing_student = (
            db.query(Student)
            .filter(Student.email == student.email)
            .first()
        )

        if existing_student:
            raise HTTPException(
                status_code=400,
                detail="Email already exists"
            )

        return StudentRepository.create_student(
            db,
            student
        )

    @staticmethod
    def get_students(
        db: Session
    ):

        return StudentRepository.get_students(db)

    @staticmethod
    def get_student(
        db: Session,
        student_id: int
    ):

        student = StudentRepository.get_student(
            db,
            student_id
        )

        if not student:
            raise HTTPException(
                status_code=404,
                detail="Student not found"
            )

        return student

    @staticmethod
    def update_student(
        db: Session,
        student_id: int,
        student: StudentUpdate
    ):

        db_student = StudentRepository.get_student(
            db,
            student_id
        )

        if not db_student:
            raise HTTPException(
                status_code=404,
                detail="Student not found"
            )

        duplicate = (
            db.query(Student)
            .filter(
                Student.email == student.email,
                Student.id != student_id
            )
            .first()
        )

        if duplicate:
            raise HTTPException(
                status_code=400,
                detail="Email already exists"
            )

        return StudentRepository.update_student(
            db,
            db_student,
            student
        )

    @staticmethod
    def delete_student(
        db: Session,
        student_id: int
    ):

        db_student = StudentRepository.get_student(
            db,
            student_id
        )

        if not db_student:
            raise HTTPException(
                status_code=404,
                detail="Student not found"
            )

        StudentRepository.delete_student(
            db,
            db_student
        )

        return {
            "message": "Student deleted successfully"
        }