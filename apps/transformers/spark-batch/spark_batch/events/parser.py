import asyncio
from pylog.log import setup_logging
import dto_controller.output as dto_output

class EventParser:
    def __init__(self, config: dto_output.ConfigDTO):
        self._config = config
        self._config_id = config.id

    def _get_parser(self):
        ...

    def _parse_message(self, message):
        ...

    def _validate_message(self, message):
        ...

    def parse(self, message):
        ...
