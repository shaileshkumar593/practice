from fastapi import APIRouter, Depends, HTTPException
from sqlalchemy.orm import Session

from app.database import get_db
from app.schemas.profile import (
    ProfileResponse,
    ProfileUpdate
)
from app.services.profile_service import ProfileService

router = APIRouter(
    prefix="/profiles",
    tags=["Profiles"]
)


@router.get(
    "/{profile_id}",
    response_model=ProfileResponse
)
def get_profile(
    profile_id: int,
    db: Session = Depends(get_db)
):

    profile = ProfileService.get_profile(
        db,
        profile_id
    )

    if not profile:
        raise HTTPException(
            status_code=404,
            detail="Profile not found"
        )

    return profile


@router.put(
    "/{profile_id}",
    response_model=ProfileResponse
)
def update_profile(
    profile_id: int,
    profile: ProfileUpdate,
    db: Session = Depends(get_db)
):

    db_profile = ProfileService.get_profile(
        db,
        profile_id
    )

    if not db_profile:
        raise HTTPException(
            status_code=404,
            detail="Profile not found"
        )

    db.commit()

    db.refresh(db_profile)

    return ProfileService.update_profile(
        db,
        db_profile,
        profile
    )


@router.delete("/{profile_id}")
def delete_profile(
    profile_id: int,
    db: Session = Depends(get_db)
):

    db_profile = ProfileService.get_profile(
        db,
        profile_id
    )

    if not db_profile:
        raise HTTPException(
            status_code=404,
            detail="Profile not found"
        )

    ProfileService.delete_profile(
        db,
        db_profile
    )

    db.commit()

    return {
        "message": "Profile deleted successfully"
    }