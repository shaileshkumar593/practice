import asyncio

async def worker(n):

    await asyncio.sleep(2)

    print(n)

async def main():

    tasks = []

    for i in range(5):

        tasks.append(
            asyncio.create_task(worker(i))
        )

    await asyncio.gather(*tasks)

asyncio.run(main())