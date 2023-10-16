from dataclasses import dataclass, field
from datetime import datetime
from typing import List
import os
from pyspark import SparkConf
from pyspark.sql import SparkSession
from pyspark.sql.types import StructType, StructField, StringType, IntegerType
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
        self._context = config.context
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


    def _get_bucket_name(self):
        return "s3a://bronze-{context}-source-{source}".format(
            context=self._context,
            source=self._source_name,
        )

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
        logger.info("[SPARK-CONNECT] Creating REMOTE Spark Session...")
        bucket_name = self._get_bucket_name()
        spark_session = (
            SparkSession
            .builder
            .remote("sc://spark:15002")
            # .config("spark.sql.catalogImplementation", "hive")
            # .config("spark.hive.metastore.uris", "thrift://hive-metastore:9083")
            .config('spark.sql.catalog.iceberg', 'org.apache.iceberg.spark.SparkCatalog')
            .config('spark.sql.catalog.iceberg.type', 'hadoop')
            # .config('spark.sql.catalog.iceberg.uri', 'thrift://hive-metastore:9083')
            .config('spark.sql.catalog.iceberg.warehouse', 's3a://iceberg-warehouse')
            .config('spark.sql.catalog.iceberg.s3.path-style-access', 'true')
            .config("spark.sql.catalog.iceberg.io-impl", "org.apache.iceberg.aws.s3.S3FileIO")
            .config("spark.hadoop.fs.s3a.impl", "org.apache.hadoop.fs.s3a.S3AFileSystem")
            .config("spark.sql.catalog.spark_catalog", "org.apache.iceberg.spark.SparkSessionCatalog")
            .config("spark.sql.catalog.spark_catalog.type", "hadoop")
            .config("spark.hadoop.fs.s3a.region", "us-east-1")


            # # .config("spark.jars.packages","org.apache.spark:spark-connect_2.12:3.4.0,org.apache.iceberg:iceberg-spark-runtime-3.2_2.12/1.3.1,org.projectnessie.nessie-integrations:nessie-spark-extensions-3.1_2.12/0.71.1,software.amazon.awssdk/bundle/2.20.156,software.amazon.awssdk/url-connection-client/2.20.156")
            # # .config("spark.jars.packages","org.apache.spark:spark-connect_2.12:3.4.1,org.apache.iceberg:iceberg-spark-runtime-3.4_2.12:1.3.0,org.projectnessie.nessie-integrations:nessie-spark-extensions-3.4_2.12:0.71.1,software.amazon.awssdk:bundle:2.17.178,software.amazon.awssdk:url-connection-client:2.17.178")
            # # .config("spark.sql.extensions", "org.apache.iceberg.spark.extensions.IcebergSparkSessionExtensions,org.projectnessie.spark.extensions.NessieSparkSessionExtensions")
            # .config("spark.sql.catalog.nessie", "org.projectnessie.spark.extensions.NessieCatalog")
            # .config("spark.sql.catalog.nessie.uri", "http://nessie:19120/api/v1")
            # .config("spark.sql.catalog.nessie.ref", "main")
            # .config("spark.sql.catalog.nessie.warehouse", bucket_name)
            # .config("spark.sql.catalog.nessie.authentication.type", "NONE")
            # .config("spark.sql.catalog.nessie.catalog-impl", "org.apache.iceberg.nessie.NessieCatalog")
            # .config("spark.sql.catalog.iceberg.s3.endpoint", "http://minio:9000/")
            # # .config("spark.sql.catalog.nessie.io-impl", "org.apache.iceberg.aws.s3.S3FileIO")
            # .config("spark.hadoop.fs.s3a.connection.ssl.enabled", "false")
            # .config("spark.hadoop.fs.s3a.path.style.access", "true")
            # .config("spark.hadoop.fs.s3a.fast.upload", "true")
            # .config("spark.hadoop.fs.s3a.access.key", "k7fK4ZHMhXtzammjr34B")
            # .config("spark.hadoop.fs.s3a.secret.key", "LyHAtFNFGkOf6bnhEH69LUeEXXLjyRHlacSrODWZ")
            .config("log4j.logger.org.apache.hadoop.metrics2", "WARN")
            .getOrCreate()
        )
        logger.info("[SPARK-CONNECT] Spark Running...")

        df = spark_session.read.format("s3selectCSV").csv(f"{uri}/part-1.csv", header=True, inferSchema=True)
        logger.info(f"Dataframe schema: {df}")
        ## Create a Table

        ## Run a Query to create a table
        spark_session.sql("CREATE TABLE iceberg.table1 (name string) USING iceberg;")

        ## Run a Query to insert into the table
        spark_session.sql("INSERT INTO iceberg.table1 VALUES ('Alex'), ('Dipankar'), ('Jason')")

        ## Run a Query to get data
        df = spark_session.sql("SELECT * FROM iceberg.table1")

        ## Show the data
        logger.info(f"Dataframe: {df}")
        df.show()

        # minio = minio_client()
        # uri = minio.upload_bytes(self._get_bucket_name(), f"{self._patition}/{self._source}.zip", response.content)
        logger.info(f"File storage uri: {uri}")
        result = {"documentUri": uri, "partition": partition}
        logger.info(f"Job result: {result}")
        spark_session.stop()
        return result, self._get_status(), uri
