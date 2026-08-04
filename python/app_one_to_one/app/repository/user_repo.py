from sqlalchemy.orm import Session

from app.models.user import User
from app.schemas.user import UserCreate, UserUpdate

from app.repository.profile_repository import ProfileRepository


class UserRepository:

    @staticmethod
    def create_user(
        db: Session,
        user: UserCreate
    ) -> User:

        db_user = User(
            name=user.name,
            email=user.email
        )

        db.add(db_user)

        db.flush()

        ProfileRepository.create_profile(
            db,
            user.profile,
            db_user.id
        )

        db.commit()

        db.refresh(db_user)

        return db_user


    @staticmethod
    def get_all_users(
        db: Session
    ):

        return db.query(User).all()


    @staticmethod
    def get_user(
        db: Session,
        user_id: int
    ):

        return db.query(User).filter(
            User.id == user_id
        ).first()


    @staticmethod
    def update_user(
        db: Session,
        db_user: User,
        user: UserUpdate
    ):

        db_user.name = user.name
        db_user.email = user.email

        db_user.profile.age = user.profile.age
        db_user.profile.city = user.profile.city
        db_user.profile.phone = user.profile.phone

        db.commit()

        db.refresh(db_user)

        return db_user


    @staticmethod
    def delete_user(
        db: Session,
        db_user: User
    ):

        db.delete(db_user)

        db.commit()