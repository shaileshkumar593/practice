import logging, time
from starlette.middleware.base import BaseHTTPMiddleware
logger=logging.getLogger("api")
class AccessLogMiddleware(BaseHTTPMiddleware):
    async def dispatch(self, request, call_next):
        started=time.perf_counter(); response=await call_next(request)
        logger.info("%s %s status=%s latency_ms=%.2f",request.method,request.url.path,response.status_code,(time.perf_counter()-started)*1000)
        return response
