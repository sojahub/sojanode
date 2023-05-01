import datetime
from typing import Tuple
from sojatool.common import *
from sojatool import cosmos, sojahub


def get_block_times(sojanoded: sojahub.Sojanoded, first_block: int, last_block: int) -> List[Tuple[int, datetime.datetime]]:
    result = [(block, cosmos.parse_iso_timestamp(sojanoded.query_block(block)["block"]["header"]["time"]))
        for block in range(first_block, last_block)]
    return result
