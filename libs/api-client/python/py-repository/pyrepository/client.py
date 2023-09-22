from pyrequest.factory import RateLimitedAsyncHttpClient
from pysd.service_discovery import new_from_env
from dto_repository.output import SchemaDTO
from serializer.serialize import from_data_to_DTO


class AsyncPyRepositoryClient:
    def __init__(self, base_url):
        self.__max_calls = 100
        self.__period = 60
        self.client = RateLimitedAsyncHttpClient(base_url, self.__max_calls, self.__period)

    async def list_one_schema_by_service_source_and_schema_type(self, service_name: str, source_name: str, schema_type: str) -> SchemaDTO:
        endpoint = "/schemas/service/{service_name}/source/{source_name}/schema-type/{schema_type}".format(
            service_name=service_name,
            source_name=source_name,
            schema_type=schema_type
        )
        result = await self.client.make_request("GET", endpoint)
        return from_data_to_DTO(result, SchemaDTO)


def async_pyrepository_client():
    sd = new_from_env()
    return AsyncPyRepositoryClient(sd.lake_repository_endpoint())
