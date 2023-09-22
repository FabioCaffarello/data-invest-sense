import asyncio
import os
from typing import List

from config_loader.loader import fetch_configs
from consumer.consumer import EventConsumer
from pylog.log import setup_logging
from pyrabbit.consumer import RabbitMQConsumer
from pysd.service_discovery import ServiceDiscovery, new_from_env

logger = setup_logging(__name__, log_level="DEBUG")


QUEUE_ACTIVE_JOBS = asyncio.Queue()
SERVICE_NAME = os.getenv("SERVICE_NAME")
CONTEXT_ENV = os.getenv("CONTEXT_ENV")


async def create_consumers_channel(sd: ServiceDiscovery) -> List[asyncio.Task]:
    configs = await fetch_configs(SERVICE_NAME, CONTEXT_ENV)
    rabbitmq_service = RabbitMQConsumer(url=sd.rabbitmq_endpoint())
    await rabbitmq_service.connect()
    tasks = list()

    for _, context_configs in configs.items():
        for _, config in context_configs.items():
            logger.info(f"Creating consumer for config: {config.id}")
            tasks.append(
                asyncio.create_task(
                    EventConsumer(sd, rabbitmq_service, config, QUEUE_ACTIVE_JOBS).run()
                )
            )
    return tasks


async def main():
    logger.info(f"Starting {SERVICE_NAME} service")

    sd = new_from_env()
    tasks = await create_consumers_channel(sd)

    await asyncio.gather(*tasks)


if __name__ == "__main__":
    asyncio.run(main())
