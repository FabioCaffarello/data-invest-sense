from dataclasses import dataclass, fields
from pylog.log import setup_logging

logger = setup_logging(__name__)

# def generate_dataclass_from_json_schema(schema, class_name=None):
#     @dataclass
#     class DataClass:
#         pass

#     if class_name is not None:
#         DataClass.__name__ = class_name

#     for prop, details in schema.get("properties", {}).items():
#         logger.info(f"Property: {prop}")
#         data_type = details.get("type", "string")
#         default = details.get("default", None)
#         required = prop in schema.get("required", [])

#         if required:
#             setattr(DataClass, prop, None)
#         elif default is not None:
#             setattr(DataClass, prop, default)
#         else:
#             if data_type == "object":
#                 nested_class_name = prop.capitalize()
#                 nested_class = generate_dataclass_from_json_schema(details, class_name=nested_class_name)
#                 setattr(DataClass, prop, nested_class())
#             elif data_type == "array":
#                 item_type = details["items"]["type"]
#                 if item_type == "object":
#                     item_schema = details["items"]
#                     nested_class_name = prop.capitalize()
#                     nested_class = generate_dataclass_from_json_schema(item_schema, class_name=nested_class_name)
#                     setattr(DataClass, prop, [nested_class()] if not required else [])
#                 else:
#                     default = details["items"].get("default", [])
#                     setattr(DataClass, prop, default if not required else [])
#             elif default is None:
#                 setattr(DataClass, prop, None)


#     # Set metadata for fields with "json_name" attribute
#     for field_obj in fields(DataClass):
#         field_name = field_obj.name
#         setattr(getattr(DataClass, field_name), 'metadata', {"json_name": field_name})

#     logger.info(f"DataClass: {DataClass}")
#     logger.info(f"DataClass fields: {fields(DataClass)}")

#     return DataClass








def generate_dataclass_from_json_schema(schema, class_name=None):
    logger.info(f"Schema: {schema}")
    @dataclass
    class DataClass:
        pass

    if class_name is not None:
        DataClass.__name__ = class_name

    for prop, details in schema.get("properties", {}).items():
        logger.info(f"Property: {prop}")
        data_type = details.get("type", "string")
        required = prop in schema.get("required", [])
        default_value = details.get("default", None)

        if data_type == "object":
            nested_class_name = prop.capitalize()
            nested_class = generate_dataclass_from_json_schema(details, class_name=nested_class_name)
            setattr(DataClass, prop, nested_class())
        elif data_type == "array":
            item_type = details["items"]["type"]
            if item_type == "object":
                item_schema = details["items"]
                nested_class_name = prop.capitalize()
                nested_class = generate_dataclass_from_json_schema(item_schema, class_name=nested_class_name)
                setattr(DataClass, prop, [nested_class()] if default_value is None else default_value)
            else:
                setattr(DataClass, prop, default_value if default_value is not None else [])
        elif default_value is not None:
            setattr(DataClass, prop, default_value)
        else:
            setattr(DataClass, prop, None)

    # Set metadata for fields with "json_name" attribute
    for field_obj in fields(DataClass):
        field_name = field_obj.name
        setattr(getattr(DataClass, field_name), 'metadata', {"json_name": field_name})

    logger.info(f"DataClass: {DataClass}")
    logger.info(f"DataClass fields: {list(DataClass.__annotations__)}")
    logger.info(f"DataClass fields: {fields(DataClass)}")

    return DataClass

# from dataclasses import dataclass, field, asdict
# import json

# def generate_dataclass_from_json_schema(schema, class_name=None):
#     logger.info(f"Schema: {schema}")

#     @dataclass
#     class DataClass:
#         pass

#     if class_name is not None:
#         DataClass.__name__ = class_name

#     for prop, details in schema.get("properties", {}).items():
#         logger.info(f"Property: {prop}")
#         data_type = details.get("type", "string")
#         required = prop in schema.get("required", [])
#         default_value = details.get("default", None)

#         # Define a field for each property in the data class
#         if data_type == "object":
#             nested_class_name = prop.capitalize()
#             nested_class = generate_dataclass_from_json_schema(details, class_name=nested_class_name)
#             setattr(DataClass, prop, field(default_factory=lambda: asdict(nested_class())))
#         elif data_type == "array":
#             item_type = details["items"]["type"]
#             if item_type == "object":
#                 item_schema = details["items"]
#                 nested_class_name = prop.capitalize()
#                 nested_class = generate_dataclass_from_json_schema(item_schema, class_name=nested_class_name)
#                 setattr(DataClass, prop, field(default_factory=lambda: [asdict(nested_class())]))
#             else:
#                 setattr(DataClass, prop, default_value if default_value is not None else [])
#         elif default_value is not None:
#             setattr(DataClass, prop, field(default=default_value))
#         else:
#             setattr(DataClass, prop, None)

#     # Set metadata for fields with "json_name" attribute
#     for field_obj in DataClass.__dataclass_fields__.values():
#         field_name = field_obj.name
#         setattr(field_obj, 'metadata', {"json_name": field_name})

#     logger.info(f"DataClass: {DataClass}")
#     logger.info(f"DataClass fields: {list(DataClass.__annotations__)}")

#     return DataClass
