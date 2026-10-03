from fastapi import APIRouter
router=APIRouter(prefix="/files",tags=["files"])
@router.get("/static-example")
async def static_example(): return {"message":"Static files are mounted from /static"}
