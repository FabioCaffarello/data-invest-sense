from dataclasses import dataclass, field

@dataclass
class Metadata:
    domain_name: str = field(metadata={"json_name": "domain_name"})
    domain_description: str = field(metadata={"json_name": "domain_description"})
    domain_location: str = field(metadata={"json_name": "domain_location"})
    domain_type: str = field(metadata={"json_name": "domain_type"})
    domain_code_reference: str = field(metadata={"json_name": "domain_code_reference"})
    quality_rules: ...

@dataclass
class Domain:
    metadata: ... = ...
    domain_input: ... = ...
    domain_output: ... = ...
    rules: ... = ...
