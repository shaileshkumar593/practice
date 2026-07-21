from fastapi import APIRouter
router=APIRouter(prefix="/borrow",tags=["Borrow"])

@router.get("/")
def list_borrowings():
    return []
