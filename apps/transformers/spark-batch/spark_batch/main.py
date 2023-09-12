import asyncio

from pylog.log import setup_logging
from pyrabbit.consumer import RabbitMQConsumer
from pysd.service_discovery import new_from_env

logger = setup_logging(__name__, log_level="DEBUG")

async def 
