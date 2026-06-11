import asyncio

sem = asyncio.Semaphore(3)

async def worker(num):

    async with sem:

        print(f"Start {num}")

        await asyncio.sleep(2)

        print(f"End {num}")

async def main():

    await asyncio.gather(
        *[worker(i) for i in range(10)]
    )

asyncio.run(main())