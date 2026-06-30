from sqlalchemy import select, func

from app.models.order import Order


class OrderRepository:

    def create(
        self,
        db,
        order,
    ):
        db.add(order)
        db.commit()
        db.refresh(order)

        return order

    def get_all(
        self,
        db,
        page,
        size,
    ):

        query = select(Order)

        total = db.scalar(
            select(func.count()).select_from(Order)
        )

        orders = (
            db.execute(
                query.offset((page - 1) * size).limit(size)
            )
            .scalars()
            .all()
        )

        return orders, total