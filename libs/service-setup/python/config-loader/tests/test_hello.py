"""Hello unit test module."""

from config_loader.hello import hello


def test_hello():
    """Test the hello function."""
    assert hello() == "Hello service-setup-python-config-loader"
