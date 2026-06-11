import asyncio

queue = asyncio.Queue()

async def producer():

    for i in range(5):

        await queue.put(i)

        print(f"Produced {i}")

async def consumer():

    while True:

        item = await queue.get()

        print(f"Consumed {item}")

        queue.task_done()

async def main():

    asyncio.create_task(consumer())

    await producer()

    await queue.join()

asyncio.run(main())