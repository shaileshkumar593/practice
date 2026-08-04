from fastapi import HTTPException
from sqlalchemy.orm import Session

from app.models.user import User
from app.repository.user_repository import UserRepository
from app.schemas.user import UserCreate, UserUpdate


class UserService:

    @staticmethod
    def create_user(
        db: Session,
        user: UserCreate
    ):

        existing_user = db.query(User).filter(
            User.email == user.email
        ).first()

        if existing_user:
            raise HTTPException(
                status_code=400,
                detail="Email already exists"
            )

        return UserRepository.create_user(
            db,
            user
        )


    @staticmethod
    def get_all_users(
        db: Session
    ):

        return UserRepository.get_all_users(
            db
        )


    @staticmethod
    def get_user(
        db: Session,
        user_id: int
    ):

        user = UserRepository.get_user(
            db,
            user_id
        )

        if not user:
            raise HTTPException(
                status_code=404,
                detail="User not found"
            )

        return user


    @staticmethod
    def update_user(
        db: Session,
        user_id: int,
        user: UserUpdate
    ):

        db_user = UserRepository.get_user(
            db,
            user_id
        )

        if not db_user:
            raise HTTPException(
                status_code=404,
                detail="User not found"
            )

        return UserRepository.update_user(
            db,
            db_user,
            user
        )


    @staticmethod
    def delete_user(
        db: Session,
        user_id: int
    ):

        db_user = UserRepository.get_user(
            db,
            user_id
        )

        if not db_user:
            raise HTTPException(
                status_code=404,
                detail="User not found"
            )

        UserRepository.delete_user(
            db,
            db_user
        )

        return {
            "message": "User deleted successfully"
        }