import asyncio
import time
from collections import defaultdict, deque

class RateLimiter:

    def __init__(self, limit, window):

        self.limit = limit
        self.window = window

        self.requests = defaultdict(deque)

        self.lock = asyncio.Lock()

    async def allow(self, ip):

        async with self.lock:

            now = time.time()

            q = self.requests[ip]

            while q and q[0] <= now - self.window:
                q.popleft()

            if len(q) >= self.limit:
                return False

            q.append(now)

            return True