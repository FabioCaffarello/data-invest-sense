import json

from dataclasses import fields, is_dataclass


def _get_serialized_object(obj):
    data = {}
    for field_obj in fields(obj):
        field_name = field_obj.name
        field_metadata = field_obj.metadata
        json_name = field_metadata.get("json_name", field_name)
        field_value = getattr(obj, field_name)

        # Check if the field has the 'repr' attribute and it's True, and it's not None
        if field_value is not None and getattr(field_obj.default, 'repr', True):
            if is_dataclass(field_value):
                # If the field is another dataclass, recursively serialize it
                data[json_name] = _get_serialized_object(field_value)
            else:
                data[json_name] = field_value
    return data

def serialize_to_json(obj):
    data = _get_serialized_object(obj)
    return json.dumps(data, sort_keys=True)

def from_data_to_DTO(data, cls):
    args = {}

    for field_obj in fields(cls):
        field_name = field_obj.name
        field_metadata = field_obj.metadata
        json_name = field_metadata.get("json_name", field_name)

        if json_name in data:
            args[field_name] = data[json_name]

    return cls(**args)
