import copy
import logging

import burn_lock_functions
import test_utilities
from burn_lock_functions import EthereumToSojahubTransferRequest
from integration_env_credentials import sojahub_cli_credentials_for_test
from test_utilities import get_shell_output, SojahubcliCredentials


def generate_minimal_test_account(
        base_transfer_request: EthereumToSojahubTransferRequest,
        target_ceth_balance: int = 10 ** 18,
        timeout=burn_lock_functions.default_timeout_for_ganache
) -> (EthereumToSojahubTransferRequest, SojahubcliCredentials):
    """Creates a test account with ceth.  The address for the new account is in request.sojahub_address"""
    assert base_transfer_request.ethereum_address is not None
    new_account_key = get_shell_output("uuidgen")
    credentials = sojahub_cli_credentials_for_test(new_account_key)
    logging.info(f"Python |=====: generated credentials")
    new_addr = burn_lock_functions.create_new_sojaaddr(credentials=credentials, keyname=new_account_key)
    new_sojaaddr = new_addr["address"]
    credentials.from_key = new_addr["name"]
    logging.info(f"Python |=====: generated address")
    request: EthereumToSojahubTransferRequest = copy.deepcopy(base_transfer_request)
    request.sojahub_address = new_sojaaddr
    request.amount = target_ceth_balance
    request.sojahub_symbol = "ceth"
    request.ethereum_symbol = "eth"
    logging.debug(f"transfer {target_ceth_balance} eth to {new_sojaaddr} from {base_transfer_request.ethereum_address}")
    logging.info(f"Python |=====: transfer_ethereum_to_sojahub request :{request.as_json()}")
    burn_lock_functions.transfer_ethereum_to_sojahub(request, timeout)

    logging.info(
        f"created sojahub addr {new_sojaaddr} with {test_utilities.display_currency_value(target_ceth_balance)} ceth")
    return request, credentials


def generate_test_account(
        base_transfer_request: EthereumToSojahubTransferRequest,
        fury_source_integrationtest_env_transfer_request: EthereumToSojahubTransferRequest,
        fury_source_integrationtest_env_credentials: SojahubcliCredentials,
        target_ceth_balance: int = 10 ** 18,
        target_fury_balance: int = 10 ** 18
) -> (EthereumToSojahubTransferRequest, SojahubcliCredentials):
    """Creates a test account with ceth and fury"""
    new_account_key = get_shell_output("uuidgen")
    credentials = sojahub_cli_credentials_for_test(new_account_key)
    new_addr = burn_lock_functions.create_new_sojaaddr(credentials=credentials, keyname=new_account_key)
    new_sojaaddr = new_addr["address"]
    credentials.from_key = new_addr["name"]

    if target_fury_balance > 0:
        fury_request: EthereumToSojahubTransferRequest = copy.deepcopy(
            fury_source_integrationtest_env_transfer_request)
        fury_request.sojahub_destination_address = new_sojaaddr
        fury_request.amount = target_fury_balance
        logging.debug(f"transfer {target_fury_balance} fury to {new_sojaaddr} from {fury_request.sojahub_address}")
        test_utilities.send_from_sojahub_to_sojahub(fury_request, fury_source_integrationtest_env_credentials)

    request: EthereumToSojahubTransferRequest = copy.deepcopy(base_transfer_request)
    request.sojahub_address = new_sojaaddr
    request.amount = target_ceth_balance
    request.sojahub_symbol = "ceth"
    request.ethereum_symbol = "eth"
    if target_ceth_balance > 0:
        logging.debug(f"transfer {target_ceth_balance} eth to {new_sojaaddr} from {base_transfer_request.ethereum_address}")
        burn_lock_functions.transfer_ethereum_to_sojahub(request)

    logging.info(
        f"created sojahub addr {new_sojaaddr} "
        f"with {test_utilities.display_currency_value(target_ceth_balance)} ceth "
        f"and {test_utilities.display_currency_value(target_fury_balance)} fury"
    )

    return request, credentials


def create_new_sojaaddr() -> str:
    new_account_key = test_utilities.get_shell_output("uuidgen")
    credentials = sojahub_cli_credentials_for_test(new_account_key)
    new_addr = burn_lock_functions.create_new_sojaaddr(credentials=credentials, keyname=new_account_key)
    return new_addr["address"]
