from fastapi import APIRouter, Depends
from sqlalchemy.orm import Session

from app.database import get_db

from app.schemas.address import (
    AddressCreate,
    AddressUpdate,
    AddressResponse
)

from app.services.address_service import AddressService
from app.repository.address_repository import AddressRepository

router = APIRouter(
    tags=["Addresses"]
)


@router.post(
    "/students/{student_id}/addresses",
    response_model=list[AddressResponse],
    status_code=201
)
def add_addresses(
    student_id: int,
    addresses: list[AddressCreate],
    db: Session = Depends(get_db)
):
    return AddressRepository.add_addresses(
        db,
        student_id,
        addresses
    )


@router.get(
    "/students/{student_id}/addresses",
    response_model=list[AddressResponse]
)
def get_student_addresses(
    student_id: int,
    db: Session = Depends(get_db)
):
    return AddressService.get_student_addresses(
        db,
        student_id
    )


@router.get(
    "/addresses/{address_id}",
    response_model=AddressResponse
)
def get_address(
    address_id: int,
    db: Session = Depends(get_db)
):
    return AddressService.get_address(
        db,
        address_id
    )


@router.put(
    "/addresses/{address_id}",
    response_model=AddressResponse
)
def update_address(
    address_id: int,
    address: AddressUpdate,
    db: Session = Depends(get_db)
):
    return AddressService.update_address(
        db,
        address_id,
        address
    )


@router.delete("/addresses/{address_id}")
def delete_address(
    address_id: int,
    db: Session = Depends(get_db)
):
    return AddressService.delete_address(
        db,
        address_id
    )