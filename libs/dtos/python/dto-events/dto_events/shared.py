from dataclasses import dataclass, field
from typing import Dict


@dataclass
class StatusDTO:
    code: int = field(metadata={"json_name": "code"})
    detail: str = field(metadata={"json_name": "detail"})


@dataclass
class MetadataInputOriginDTO:
    gateway: str = field(metadata={"json_name": "gateway"})
    controller: str = field(metadata={"json_name": "controller"})


@dataclass
class MetadataInputDTO:
    id : str = field(metadata={"json_name": "id"})
    data: Dict[str, any] = field(metadata={"json_name": "data"})
    processing_id: str = field(metadata={"json_name": "processing_id"})
    processing_timestamp: str = field(metadata={"json_name": "processing_timestamp"})
    source: MetadataInputOriginDTO = field(metadata={"json_name": "source"})


@dataclass
class MetadataDTO:
    input : MetadataInputDTO = field(metadata={"json_name": "input"})
    service: MetadataInputOriginDTO = field(metadata={"json_name": "service"})
    context: str = field(metadata={"json_name": "context"})
    processing_id: str = field(metadata={"json_name": "processing_id"})
    processing_timestamp: str = field(metadata={"json_name": "processing_timestamp"})
    target_endpoint: str = field(metadata={"json_name": "target_endpoint"})
    job_frequency: str = field(metadata={"json_name": "job_frequency"})
