from sqlalchemy.orm import Session

from app.repository.profile_repository import ProfileRepository
from app.schemas.profile import ProfileUpdate


class ProfileService:

    @staticmethod
    def get_profile(
        db: Session,
        profile_id: int
    ):
        profile = ProfileRepository.get_profile(
            db,
            profile_id
        )

        return profile


    @staticmethod
    def update_profile(
        db: Session,
        db_profile,
        profile: ProfileUpdate
    ):

        return ProfileRepository.update_profile(
            db,
            db_profile,
            profile
        )


    @staticmethod
    def delete_profile(
        db: Session,
        db_profile
    ):

        ProfileRepository.delete_profile(
            db,
            db_profile
        )