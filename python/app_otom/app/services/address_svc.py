from sqlalchemy.orm import Session
from fastapi import HTTPException

from app.repository.address_repository import AddressRepository
from app.schemas.address import AddressUpdate


class AddressService:

    @staticmethod
    def get_address(
        db: Session,
        address_id: int
    ):

        address = AddressRepository.get_address(
            db,
            address_id
        )

        if not address:
            raise HTTPException(
                status_code=404,
                detail="Address not found"
            )

        return address

    @staticmethod
    def get_student_addresses(
        db: Session,
        student_id: int
    ):
        return AddressRepository.get_student_addresses(
            db,
            student_id
        )

    @staticmethod
    def update_address(
        db: Session,
        address_id: int,
        address: AddressUpdate
    ):

        db_address = AddressRepository.get_address(
            db,
            address_id
        )

        if not db_address:
            raise HTTPException(
                status_code=404,
                detail="Address not found"
            )

        AddressRepository.update_address(
            db,
            db_address,
            address
        )

        db.commit()

        db.refresh(db_address)

        return db_address

    @staticmethod
    def delete_address(
        db: Session,
        address_id: int
    ):

        db_address = AddressRepository.get_address(
            db,
            address_id
        )

        if not db_address:
            raise HTTPException(
                status_code=404,
                detail="Address not found"
            )

        AddressRepository.delete_address(
            db,
            db_address
        )

        db.commit()

        return {
            "message": "Address deleted successfully"
        }