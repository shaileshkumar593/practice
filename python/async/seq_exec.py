import asyncio

async def task(name):

    print(f"{name} started")

    await asyncio.sleep(2)

    print(f"{name} finished")

async def main():

    await task("A")
    await task("B")

asyncio.run(main())