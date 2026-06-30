from uuid import uuid4

from app.models.order import Order, OrderStatus, PaymentStatus
from app.repositories.order_repository import OrderRepository


class OrderService:

    def __init__(self):
        self.repo = OrderRepository()

    def create_order(self, db, request):

        total_amount = (
            request.subtotal
            - request.discount
            + request.tax
            + request.shipping_charge
        )

        order = Order(
            order_number=f"ORD-{uuid4().hex[:8].upper()}",
            customer_id=request.customer_id,

            subtotal=request.subtotal,
            discount=request.discount,
            tax=request.tax,
            shipping_charge=request.shipping_charge,
            total_amount=total_amount,

            status=OrderStatus.PENDING,
            payment_status=PaymentStatus.PENDING,
        )

        return self.repo.create(db, order)