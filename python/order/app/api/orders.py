from fastapi import (
    APIRouter,
    Depends,
    Query,
    status,
)
from sqlalchemy.orm import Session

from app.database.session import get_db
from app.schemas.order import (
    OrderCreate,
    OrderResponse,
    OrderListResponse,
)
from app.services.order_service import OrderService

router = APIRouter(
    prefix="/orders",
    tags=["Orders"],
)

service = OrderService()


@router.post(
    "",
    response_model=OrderResponse,
    status_code=status.HTTP_201_CREATED,
)
def create_order(
    request: OrderCreate,
    db: Session = Depends(get_db),
):
    return service.create_order(
        db,
        request,
    )


@router.get(
    "",
    response_model=OrderListResponse,
)
def get_orders(
    page: int = Query(1, ge=1),
    size: int = Query(10, ge=1, le=100),
    db: Session = Depends(get_db),
):
    return service.get_orders(
        db,
        page,
        size,
    )