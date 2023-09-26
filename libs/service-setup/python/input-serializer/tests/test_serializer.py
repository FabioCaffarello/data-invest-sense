import unittest
from dataclasses import dataclass, field, fields
from input_serializer.serializer import generate_dataclass_from_json_schema


class TestSerializer(unittest.TestCase):

    def test_generate_dataclass_string(self):
        schema = {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "default": "John"
                }
            }
        }
        GeneratedClass = generate_dataclass_from_json_schema(schema)
        obj = GeneratedClass()
        self.assertEqual(obj.name, "John")

    def test_generate_dataclass_integer(self):
        schema = {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer",
                    "default": 30
                }
            }
        }
        GeneratedClass = generate_dataclass_from_json_schema(schema)
        obj = GeneratedClass()
        self.assertEqual(obj.age, 30)

    def test_generate_dataclass_boolean(self):
        schema = {
            "type": "object",
            "properties": {
                "is_student": {
                    "type": "boolean",
                    "default": True
                }
            }
        }
        GeneratedClass = generate_dataclass_from_json_schema(schema)
        obj = GeneratedClass()
        self.assertTrue(obj.is_student)

    def test_generate_dataclass_number(self):
        schema = {
            "type": "object",
            "properties": {
                "price": {
                    "type": "number",
                    "default": 19.99
                }
            }
        }
        GeneratedClass = generate_dataclass_from_json_schema(schema)
        obj = GeneratedClass()
        self.assertAlmostEqual(obj.price, 19.99, places=2)

    def test_generate_dataclass_array_of_integers(self):
        schema = {
            "type": "object",
            "properties": {
                "scores": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    },
                    "default": [95, 88, 72]
                }
            }
        }
        GeneratedClass = generate_dataclass_from_json_schema(schema)
        obj = GeneratedClass()
        self.assertEqual(obj.scores, [95, 88, 72])

    def test_generate_dataclass_required_properties(self):
        schema = {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                },
                "age": {
                    "type": "integer",
                }
            },
            "required": ["name", "age"]
        }
        GeneratedClass = generate_dataclass_from_json_schema(schema)
        obj = GeneratedClass()
        # Check that required properties are initialized as None
        print("obj fields:", fields(obj))
        self.assertIsNone(obj.name)
        self.assertIsNone(obj.age)

    def test_generate_dataclass_with_nested_arrays(self):
        schema = {
            "type": "object",
            "properties": {
                "matrix": {
                    "type": "array",
                    "items": {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        }
                    },
                    "default": [[1, 2], [3, 4]]
                }
            }
        }
        GeneratedClass = generate_dataclass_from_json_schema(schema)
        obj = GeneratedClass()
        self.assertEqual(obj.matrix, [[1, 2], [3, 4]])

    def test_generate_dataclass_default_none(self):
        schema = {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "age": {
                    "type": "integer"
                }
            }
        }
        GeneratedClass = generate_dataclass_from_json_schema(schema)
        obj = GeneratedClass()
        self.assertIsNone(obj.name)
        self.assertIsNone(obj.age)

    def test_generate_dataclass_nested_object(self):
        schema = {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "default": "John"
                },
                "age": {
                    "type": "integer",
                    "default": 30
                },
                "address": {
                    "type": "object",
                    "properties": {
                        "street": {"type": "string", "default": "123 Main St"},
                        "city": {"type": "string", "default": "New York"}
                    }
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "default": []
                }
            }
        }
        GeneratedClass = generate_dataclass_from_json_schema(schema)
        obj = GeneratedClass()
        print("obj.address:", obj.address)

        self.assertEqual(obj.name, "John")
        self.assertEqual(obj.address.street, "123 Main St")
        self.assertEqual(obj.address.city, "New York")

    def test_generate_dataclass_with_required_nested_object(self):
        schema = {
            "type": "object",
            "properties": {
                "person": {
                    "type": "object",
                    "properties": {
                        "name": {"type": "string"},
                        "age": {"type": "integer"}
                    },
                    "required": ["name", "age"]
                }
            }
        }
        GeneratedClass = generate_dataclass_from_json_schema(schema)
        obj = GeneratedClass()
        # Check that required properties within nested object are initialized as None
        self.assertIsNone(obj.person.name)
        self.assertIsNone(obj.person.age)

    def test_generate_dataclass_array_of_strings(self):
        schema = {
            "type": "object",
            "properties": {
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "default": []
                }
            }
        }
        GeneratedClass = generate_dataclass_from_json_schema(schema)
        obj = GeneratedClass()
        self.assertEqual(obj.tags, [])  # Access the default value


    def test_generate_dataclass_nested_array_required_properties(self):
        schema = {
            "type": "object",
            "properties": {
                "matrix": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "value": {"type": "integer"},
                            "required_prop": {"type": "string"}
                        },
                        "required": ["required_prop"]
                    }
                }
            }
        }
        GeneratedClass = generate_dataclass_from_json_schema(schema)
        obj = GeneratedClass()
        # Check that required properties within nested array objects are initialized as None
        self.assertIsNone(obj.matrix[0].value)
        self.assertIsNone(obj.matrix[0].required_prop)

    def test_generate_dataclass_metadata(self):
        schema = {
            "properties": {
                "reference": {
                    "properties": {
                        "day": {
                            "type": "integer"
                        },
                        "month": {
                            "type": "integer"
                        },
                        "year": {
                            "type": "integer"
                        }
                    },
                    "required": [
                        "year",
                        "month",
                        "day"
                    ],
                    "type": "object"
                }
            },
            "required": [
                "reference"
            ],
            "type": "object",
            "default": {
                "year": 2017,
                "month": 1,
                "day": 1
            },
        }
        GeneratedClass = generate_dataclass_from_json_schema(schema)
        obj = GeneratedClass()

        for field_obj in fields(obj):
            field_name = field_obj.name
            field_metadata = field_obj.metadata
            json_name = field_metadata.get("json_name")
            self.assertIsNotNone(json_name)
            self.assertEqual(json_name, field_name)

if __name__ == "__main__":
    unittest.main()
