from sqlalchemy.orm import Session

from app.models.profile import Profile
from app.schemas.profile import ProfileCreate, ProfileUpdate


class ProfileRepository:

    @staticmethod
    def create_profile(
        db: Session,
        profile: ProfileCreate,
        user_id: int
    ) -> Profile:

        db_profile = Profile(
            age=profile.age,
            city=profile.city,
            phone=profile.phone,
            user_id=user_id
        )

        db.add(db_profile)

        return db_profile


    @staticmethod
    def get_profile(
        db: Session,
        profile_id: int
    ):

        return db.query(Profile).filter(
            Profile.id == profile_id
        ).first()


    @staticmethod
    def update_profile(
        db: Session,
        db_profile: Profile,
        profile: ProfileUpdate
    ):

        db_profile.age = profile.age
        db_profile.city = profile.city
        db_profile.phone = profile.phone

        return db_profile


    @staticmethod
    def delete_profile(
        db: Session,
        db_profile: Profile
    ):

        db.delete(db_profile)