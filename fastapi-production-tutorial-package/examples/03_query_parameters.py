from fastapi import FastAPI
app=FastAPI()
@app.get("/items/")
async def read_items(skip:int=0,limit:int=20): return {"skip":skip,"limit":limit}
