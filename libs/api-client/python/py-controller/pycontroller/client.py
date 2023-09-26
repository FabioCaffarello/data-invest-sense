from pyrequest.factory import RateLimitedAsyncHttpClient
from pysd.service_discovery import new_from_env
from dto_controller.output import ConfigDTO
from serializer.serialize import from_data_to_DTO

class AsyncPyControllerClient:
    def __init__(self, base_url):
        self.__max_calls = 100
        self.__period = 60
        self.client = RateLimitedAsyncHttpClient(base_url, self.__max_calls, self.__period)

    async def create_config(self, data: dict):
        endpoint = "/configs"
        config = await self.client.make_request("POST", endpoint, data=data)
        return from_data_to_DTO(config, ConfigDTO)

    async def list_all_configs(self):
        endpoint = "/configs"
        configs = await self.client.make_request("GET", endpoint)
        return [from_data_to_DTO(config, ConfigDTO) for config in configs]

    async def list_one_config_by_id(self, config_id: str):
        endpoint = f"/configs/{config_id}"
        config = await self.client.make_request("GET", endpoint)
        return from_data_to_DTO(config, ConfigDTO)

    async def list_all_configs_by_service(self, service_name: str):
        endpoint = f"/configs/service/{service_name}"
        configs = await self.client.make_request("GET", endpoint)
        return [from_data_to_DTO(config, ConfigDTO) for config in configs]

    async def list_all_configs_by_service_and_context(self, service_name: str, context: str) -> list[ConfigDTO]:
        endpoint = f"/configs/service/{service_name}/context/{context}"
        configs = await self.client.make_request("GET", endpoint)
        return [from_data_to_DTO(config, ConfigDTO) for config in configs]

def async_pycontroller_client():
    sd = new_from_env()
    return AsyncPyControllerClient(sd.lake_controller_endpoint())


