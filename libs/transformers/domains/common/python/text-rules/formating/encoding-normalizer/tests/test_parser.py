import unittest
from normalizer.parser import normalize


class TestParser(unittest.TestCase):

    def test_parse_empty_string(self):
        self.assertEqual(normalize(""), "")

    def test_parse_ascii_text(self):
        self.assertEqual(normalize("Hello, world!"), "Hello, world!")

    def test_parse_unicode_text(self):
        # Test with some common Unicode characters
        input_text = "Héllo, ẁörld! 👋🌍"
        expected_output = "Hello, world! "
        self.assertEqual(normalize(input_text), expected_output)

    def test_parse_unicode_text_with_combining_characters(self):
        # Test with combining characters
        input_text = "café"
        expected_output = "cafe"
        self.assertEqual(normalize(input_text), expected_output)

    def test_parse_unicode_text_with_non_ascii_characters(self):
        # Test with non-ASCII characters
        input_text = "Mötörhëäd"
        expected_output = "Motorhead"
        self.assertEqual(normalize(input_text), expected_output)

if __name__ == '__main__':
    unittest.main()
