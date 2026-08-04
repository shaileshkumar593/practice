from sqlalchemy import Column,Integer,String,ForeignKey
from sqlalchemy.orm import relationship

class User(Base):
    __tablename__="users"

    id=Column(Integer,primary_key=True)
    name=Column(String)

    profile=relationship(
        "Profile",
        back_populates="user",
        uselist=False 
        # tells SQLAlchemy that there is only one profile.
    )


class Profile(Base):
    __tablename__="profiles"

    id=Column(Integer,primary_key=True)
    address=Column(String)
    phone=Column(String)

    user_id=Column(
        Integer,
        ForeignKey("users.id"),
        unique=True
    )

    user=relationship(
        "User",
        back_populates="profile"
    )