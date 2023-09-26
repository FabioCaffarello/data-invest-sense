from pyrequest.factory import RateLimitedAsyncHttpClient
from pysd.service_discovery import new_from_env
from dto_lake_catalog.output import SchemaCatalogDTO
from serializer.serialize import from_data_to_DTO


class AsyncPyLakeCatalogClient:
    def __init__(self, base_url):
        self.__max_calls = 100
        self.__period = 60
        self.client = RateLimitedAsyncHttpClient(base_url, self.__max_calls, self.__period)

    async def list_one_schema_catalog_by_id(self, schema_catalog_id: str):
        endpoint = "/schemas-catalog/{schema_catalog_id}".format(
            schema_catalog_id=schema_catalog_id
        )
        result = await self.client.make_request("GET", endpoint)
        return from_data_to_DTO(result, SchemaCatalogDTO)

    async def list_one_schema_by_service_source(self, service_name: str, source_name: str):
        endpoint = "/schemas-catalog/service/{service_name}/source/{source_name}".format(
            service_name=service_name,
            source_name=source_name
        )
        result = await self.client.make_request("GET", endpoint)
        return from_data_to_DTO(result, SchemaCatalogDTO)

def async_pylakecatalog_client():
    sd = new_from_env()
    return AsyncPyLakeCatalogClient(sd.lake_catalog_endpoint())
