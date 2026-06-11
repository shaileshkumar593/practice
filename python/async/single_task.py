import asyncio

async def worker():

    await asyncio.sleep(2)

    print("Done")

async def main():

    task = asyncio.create_task(worker())

    print("Running")

    await task

asyncio.run(main())