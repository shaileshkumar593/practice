import asyncio

class Counter:

    def __init__(self):
        self.current = 0

    def __aiter__(self):
        return self

    async def __anext__(self):

        if self.current >= 5:
            raise StopAsyncIteration

        await asyncio.sleep(1)

        self.current += 1

        return self.current

async def main():

    async for item in Counter():

        print(item)

asyncio.run(main())