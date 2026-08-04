import asyncio

async def numbers():

    for i in range(5):
        await asyncio.sleep(1)
        yield i

async def main():

    squares = [
        x*x
        async for x in numbers()
    ]

    print(squares)

asyncio.run(main())