import asyncio, logging
logger=logging.getLogger(__name__)
async def write_audit_log(event: str, entity_id: int) -> None:
    await asyncio.sleep(0); logger.info("audit event=%s entity_id=%s",event,entity_id)
