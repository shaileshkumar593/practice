from sqlalchemy.orm import Session

from app.models.address import Address
from app.schemas.address import AddressCreate, AddressUpdate


class AddressRepository:

    @staticmethod
    def create_addresses(
        db: Session,
        addresses: list[AddressCreate],
        student_id: int
    ):

        db_addresses = []

        for address in addresses:

            db_address = Address(
                city=address.city,
                state=address.state,
                country=address.country,
                student_id=student_id
            )

            db.add(db_address)

            db_addresses.append(db_address)

        return db_addresses

    @staticmethod
    def get_address(
        db: Session,
        address_id: int
    ):

        return (
            db.query(Address)
            .filter(Address.id == address_id)
            .first()
        )

    @staticmethod
    def get_student_addresses(
        db: Session,
        student_id: int
    ):

        return (
            db.query(Address)
            .filter(Address.student_id == student_id)
            .all()
        )

    @staticmethod
    def update_address(
        db: Session,
        db_address: Address,
        address: AddressUpdate
    ):

        db_address.city = address.city
        db_address.state = address.state
        db_address.country = address.country

        return db_address

    @staticmethod
    def delete_address(
        db: Session,
        db_address: Address
    ):

        db.delete(db_address)

    @staticmethod
    def add_addresses(
        db: Session,
        student_id: int,
        addresses: list[AddressCreate]
    ):

        db_addresses = []

        for address in addresses:

            db_address = Address(
                city=address.city,
                state=address.state,
                country=address.country,
                student_id=student_id
            )

            db.add(db_address)

            db_addresses.append(db_address)

        db.commit()

        for address in db_addresses:
            db.refresh(address)

        return db_addresses