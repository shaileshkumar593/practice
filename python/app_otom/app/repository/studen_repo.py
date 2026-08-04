from sqlalchemy.orm import Session

from app.models.student import Student
from app.schemas.student import StudentCreate, StudentUpdate

from app.repository.address_repository import AddressRepository


class StudentRepository:

    @staticmethod
    def create_student(
        db: Session,
        student: StudentCreate
    ):

        db_student = Student(
            name=student.name,
            email=student.email
        )

        db.add(db_student)

        # Generate student.id
        db.flush()

        AddressRepository.create_addresses(
            db,
            student.addresses,
            db_student.id
        )

        db.commit()

        db.refresh(db_student)

        return db_student

    @staticmethod
    def get_students(
        db: Session
    ):

        return db.query(Student).all()

    @staticmethod
    def get_student(
        db: Session,
        student_id: int
    ):

        return (
            db.query(Student)
            .filter(Student.id == student_id)
            .first()
        )

    @staticmethod
    def update_student(
        db: Session,
        db_student: Student,
        student: StudentUpdate
    ):

        db_student.name = student.name
        db_student.email = student.email

        # Remove existing addresses
        db_student.addresses.clear()

        # Add updated addresses
        for address in student.addresses:

            db_student.addresses.append(

                AddressRepository.create_single_address(
                    address)

            )

        db.commit()

        db.refresh(db_student)

        return db_student

    @staticmethod
    def delete_student(
        db: Session,
        db_student: Student
    ):

        db.delete(db_student)

        db.commit()