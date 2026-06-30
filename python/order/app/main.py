from fastapi import FastAPI

from app.models.order import Order
from app.models.order_item import OrderItem


from app.api.orders import router as order_router

app = FastAPI(
    title="Order Service",
)

app.include_router(order_router)