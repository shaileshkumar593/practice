from datetime import datetime
from decimal import Decimal
from enum import Enum
from uuid import UUID

from pydantic import BaseModel, ConfigDict, Field


class OrderStatus(str, Enum):
    PENDING = "PENDING"
    CONFIRMED = "CONFIRMED"
    PACKED = "PACKED"
    SHIPPED = "SHIPPED"
    DELIVERED = "DELIVERED"
    CANCELLED = "CANCELLED"


class PaymentStatus(str, Enum):
    PENDING = "PENDING"
    PAID = "PAID"
    FAILED = "FAILED"
    REFUNDED = "REFUNDED"


# -----------------------------
# Order Create Request
# -----------------------------
class OrderCreate(BaseModel):
    customer_id: UUID

    subtotal: Decimal = Field(..., ge=0)
    discount: Decimal = Field(default=0, ge=0)
    tax: Decimal = Field(default=0, ge=0)
    shipping_charge: Decimal = Field(default=0, ge=0)


# -----------------------------
# Order Update Request
# -----------------------------
class OrderUpdate(BaseModel):
    status: OrderStatus | None = None
    payment_status: PaymentStatus | None = None


# -----------------------------
# Order Response
# -----------------------------
class OrderResponse(BaseModel):
    model_config = ConfigDict(from_attributes=True)

    id: UUID
    order_number: str
    customer_id: UUID

    status: OrderStatus
    payment_status: PaymentStatus

    subtotal: Decimal
    discount: Decimal
    tax: Decimal
    shipping_charge: Decimal
    total_amount: Decimal

    created_at: datetime
    updated_at: datetime


# -----------------------------
# Paginated Response
# -----------------------------
class OrderListResponse(BaseModel):
    page: int
    size: int
    total: int
    data: list[OrderResponse]