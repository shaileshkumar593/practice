from fastapi import BackgroundTasks,FastAPI
app=FastAPI()
def audit(message:str): print(message)
@app.post("/events")
async def event(background_tasks:BackgroundTasks): background_tasks.add_task(audit,"event received"); return {"accepted":True}
