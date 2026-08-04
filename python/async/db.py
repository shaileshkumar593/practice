import asyncio

class Connection:

    async def __aenter__(self):
        print("Connected")
        return self

    async def __aexit__(self, exc_type, exc, tb):
        print("Disconnected")

async def main():

    async with Connection():
        print("Running Query")

asyncio.run(main())