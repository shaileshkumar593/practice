import asyncio

async def hello():

    print("Start")

    await asyncio.sleep(2)  # time.sleep  block co routine but this doesnot

    print("End")

asyncio.run(hello())