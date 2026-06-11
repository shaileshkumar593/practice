import asyncio

async def square(n):

    await asyncio.sleep(1)

    return n * n

async def main():

    results = await asyncio.gather(
        square(2),
        square(3),
        square(4)
    )

    print(results)

asyncio.run(main())