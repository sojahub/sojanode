import copy
import logging

import pytest
import time

import burn_lock_functions
import test_utilities
from integration_env_credentials import sojahub_cli_credentials_for_test
from test_utilities import EthereumToSojahubTransferRequest, SojahubcliCredentials


@pytest.mark.skipif(not test_utilities.get_optional_env_var("DESTINATION_ACCOUNT", None), reason="run by hand and specify DESTINATION_ACCOUNT")
def test_token_distribution(
        basic_transfer_request: EthereumToSojahubTransferRequest,
        fury_source_integrationtest_env_credentials: SojahubcliCredentials,
        fury_source_integrationtest_env_transfer_request: EthereumToSojahubTransferRequest,
        smart_contracts_dir,
        source_ethereum_address,
        fury_source,
        fury_source_key,
        bridgebank_address,
        bridgetoken_address,
        ethereum_network,
):
    tokens = test_utilities.get_whitelisted_tokens(basic_transfer_request)
    request = basic_transfer_request
    amount_in_tokens = 10000000

    for t in tokens:
        try:
            logging.info(f"sending: {t}")
            destination_symbol = "c" + t["symbol"]
            if t["symbol"] == "efury":
                destination_symbol = "fury"
            request.amount = int(amount_in_tokens * (10 ** int(t["decimals"])))
            request.ethereum_symbol = t["token"]
            request.sojahub_symbol = destination_symbol
            request.sojahub_address = fury_source
            request.sojahub_destination_address = test_utilities.get_required_env_var("DESTINATION_ACCOUNT")
            test_utilities.send_from_sojahub_to_sojahub(request, fury_source_integrationtest_env_credentials)
        except Exception as e:
            logging.error(f"error: {e}")
