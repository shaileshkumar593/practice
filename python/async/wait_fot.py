import asyncio

async def task1():
    await asyncio.sleep(1)
    return "task1"

async def task2():
    await asyncio.sleep(5)
    return "task2"

async def main():

    t1 = asyncio.create_task(task1())
    t2 = asyncio.create_task(task2())

    done, pending = await asyncio.wait(
        [t1, t2],
        timeout=2
    )

    print("Done:", done)
    print("Pending:", pending)

asyncio.run(main())