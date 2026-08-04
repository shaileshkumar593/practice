import asyncio

async def numbers():

    for i in range(5):

        await asyncio.sleep(1)

        yield i

async def main():

    async for number in numbers():

        print(number)

asyncio.run(main())