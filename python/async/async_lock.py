import asyncio

counter = 0

lock = asyncio.Lock()

async def worker():

    global counter

    for _ in range(1000):

        async with lock:

            counter += 1

async def main():

    await asyncio.gather(
        worker(),
        worker(),
        worker()
    )

    print(counter)

asyncio.run(main())