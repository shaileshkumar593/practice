from sqlalchemy import Column
from sqlalchemy import ForeignKey
from sqlalchemy import Integer
from sqlalchemy import String

from sqlalchemy.orm import relationship

from app.database import Base


class Address(Base):

    __tablename__ = "addresses"

    id = Column(
        Integer,
        primary_key=True,
        index=True
    )

    city = Column(
        String(100),
        nullable=False
    )

    state = Column(
        String(100),
        nullable=False
    )

    country = Column(
        String(100),
        nullable=False
    )

    student_id = Column(
        Integer,
        ForeignKey("students.id"),
        nullable=False
    )

    student = relationship(
        "Student",
        back_populates="addresses"
    )