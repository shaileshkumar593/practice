import asyncio

async def numbers():
    for i in range(5):
        await asyncio.sleep(1)
        yield i

async def main():

    squares = (x * x async for x in numbers())

    async for value in squares:
        print(value)

asyncio.run(main())


"""
Nothing executes immediately.

Execution begins when iterated.
"""