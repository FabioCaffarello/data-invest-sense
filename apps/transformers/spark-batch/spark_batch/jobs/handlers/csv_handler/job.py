from dataclasses import dataclass, field
from datetime import datetime
from typing import List
from pyspark.sql import SparkSession
import requests
from dto_controller.output import ConfigDTO
from dto_events.shared import StatusDTO
from serializer.serialize import from_data_to_DTO
from pylog.log import setup_logging
from pyminio.pyminio import minio_client
from pylakecatalog.client import async_pylakecatalog_client

logger = setup_logging(__name__)

@dataclass
class FieldsDTO:
    file_names: List[str] = field(metadata={"json_name": "file_names"})
    job_name: str = field(metadata={"json_name": "job_name"})

@dataclass
class MappingHeaderDTO:
    properties: List[FieldsDTO] = field(metadata={"json_name": "properties"})


class Job:
    def __init__(self, config: ConfigDTO, input_data):
        self._config = config
        self._config_id = config.id
        self._service_name = config.service
        self._source_name = config.source
        self._input_data = input_data
        self._lake_catalog_client = async_pylakecatalog_client()


    async def _get_mapping_headers_from_catalog_metadata(self):
        job_catalog = await self._lake_catalog_client.list_one_schema_by_service_source(
            service_name=self._service_name,
            source_name=self._source_name
        )
        mapping_header = job_catalog.catalog
        return from_data_to_DTO(mapping_header, MappingHeaderDTO)


    # def _get_bucket_name(self):
    #     return "landing-{context}-source-{source}".format(
    #         context=self._context,
    #         source=self._source,
    #     )

    def _get_status(self) -> StatusDTO:
        return StatusDTO(
            code=200,
            detail="OK",
        )


    async def run(self):
        mapping_headers = await self._get_mapping_headers_from_catalog_metadata()
        logger.info(f"Mapping headers: {mapping_headers}")
        uri = self._input_data.documentUri
        logger.info(f"Input data uri: {uri}")
        partition = self._input_data.partition
        logger.info(f"Input data partition: {partition}")
        spark_session = SparkSession.builder.remote("sc://localhost").getOrCreate()
        logger.info("Spark session created")
        df = spark_session.read.csv(f"{uri}/*.csv", header=True, inferSchema=True)
        logger.info(f"Dataframe schema: {df}")


        # minio = minio_client()
        # uri = minio.upload_bytes(self._get_bucket_name(), f"{self._patition}/{self._source}.zip", response.content)
        logger.info(f"File storage uri: {uri}")
        result = {"documentUri": uri, "partition": partition}
        logger.info(f"Job result: {result}")
        return result, self._get_status(), uri
