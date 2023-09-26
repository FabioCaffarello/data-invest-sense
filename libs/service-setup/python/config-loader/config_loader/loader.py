
from typing import Dict
from pylog.log import setup_logging
from pycontroller.client import async_pycontroller_client
from dto_controller.output import ConfigDTO

logger = setup_logging(__name__)


mapping_config: Dict[str, Dict[str, ConfigDTO]] = dict()

async def fetch_configs(service, context_env):
    await ConfigLoader().fetch_configs_for_service(service_name=service, context_env=context_env)
    return mapping_config


class ConfigLoader:
    def __init__(self) -> None:
        self.__controller_api_client = async_pycontroller_client()
        super().__init__()

    async def fetch_configs_for_service(self, service_name: str, context_env: str):
        configs = await self.__controller_api_client.list_all_configs_by_service_and_context(service_name, context_env)
        for config in configs:
            register_config(
                config.context,
                config.id,
                config
            )


def register_config(context: str, config_id: str, config: ConfigDTO):
    if context not in mapping_config:
        mapping_config[context] = dict()
    if config_id in mapping_config[context]:
        raise Exception(f"Duplicate config ID '{config_id}' for context '{context}'. Overwriting existing config.")
    mapping_config[context][config_id] = config
