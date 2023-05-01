import copy
import logging
import os
import threading

import pytest

import test_utilities
from burn_lock_functions import decrease_log_level, force_log_level


@pytest.fixture
def smart_contracts_dir(sojanode_base_dir):
    return test_utilities.get_optional_env_var("SMART_CONTRACTS_DIR", os.path.join(sojanode_base_dir, "smart-contracts"))


@pytest.fixture
def sojanode_base_dir():
    return test_utilities.get_required_env_var("BASEDIR")


@pytest.fixture
def sojahub_admin_account():
    return test_utilities.get_required_env_var("SOJAHUB_ADMIN_ACCOUNT")


@pytest.fixture
def sojahub_admin_account_credentials(sojahub_admin_account):
    return test_utilities.SojahubcliCredentials(
        from_key=sojahub_admin_account
    )


@pytest.fixture
def smart_contract_artifact_dir(smart_contracts_dir):
    result = test_utilities.get_optional_env_var("SMART_CONTRACT_ARTIFACT_DIR", None)
    return result if result else os.path.join(smart_contracts_dir, "build/contracts")


@pytest.fixture
def validator_password():
    return test_utilities.get_optional_env_var("VALIDATOR1_PASSWORD", None)


@pytest.fixture
def validator_address():
    return test_utilities.get_optional_env_var("VALIDATOR1_ADDR", None)


@pytest.fixture
def integration_dir():
    return test_utilities.get_required_env_var("TEST_INTEGRATION_DIR")


@pytest.fixture
def bridgebank_address(smart_contract_artifact_dir, ethereum_network_id):
    return env_or_truffle_artifact("BridgeBank", "BRIDGE_BANK_ADDRESS", smart_contract_artifact_dir,
                                   ethereum_network_id)


def env_or_truffle_artifact(contract_name, contract_env_var, smart_contract_artifact_dir, ethereum_network_id):
    result = test_utilities.get_optional_env_var(contract_env_var, None)
    return result if result else test_utilities.contract_address(
        smart_contract_artifact_dir=smart_contract_artifact_dir,
        contract_name=contract_name,
        ethereum_network_id=ethereum_network_id
    )


@pytest.fixture
def bridgetoken_address(smart_contract_artifact_dir, ethereum_network_id):
    return env_or_truffle_artifact("BridgeToken", "BRIDGE_TOKEN_ADDRESS", smart_contract_artifact_dir,
                                   ethereum_network_id)


@pytest.fixture
def ethereum_network():
    return test_utilities.get_optional_env_var("ETHEREUM_NETWORK", "")


@pytest.fixture
def n_sojahub_accounts():
    return int(test_utilities.get_optional_env_var("N_SOJAHUB_ACCOUNTS", 1))


@pytest.fixture
def fury_amount():
    """the meaning of fury_amount is determined by the test using it"""
    return int(int(test_utilities.get_optional_env_var("FURY_AMOUNT", 10 ** 18)))


@pytest.fixture
def solidity_json_path(smart_contracts_dir):
    return test_utilities.get_optional_env_var("SOLIDITY_JSON_PATH", f"{smart_contracts_dir}/build/contracts")


@pytest.fixture
def sojanoded_homedir(is_ropsten_testnet):
    if is_ropsten_testnet:
        base = test_utilities.get_required_env_var("HOME")
    else:
        base = test_utilities.get_required_env_var("CHAINDIR")
    result = f"""{base}/.sojanoded"""
    return result


@pytest.fixture
def fury_source(is_ropsten_testnet, validator_address):
    """A sojahub address or key that has fury and can send that fury to other address"""
    result = test_utilities.get_optional_env_var("FURY_SOURCE", None)
    if result:
        return result
    if is_ropsten_testnet:
        assert result
    else:
        assert validator_address
        return validator_address


@pytest.fixture
def fury_source_key(is_ropsten_testnet, fury_source):
    """A sojahub address or key that has fury and can send that fury to other address"""
    result = test_utilities.get_optional_env_var("FURY_SOURCE_KEY", fury_source)
    if result:
        return result
    if is_ropsten_testnet:
        # Ropsten requires that you manually set the FURY_SOURCE_KEY environment variable
        assert result
    else:
        return test_utilities.get_required_env_var("MONIKER")


@pytest.fixture
def sojanoded_node():
    return test_utilities.get_optional_env_var("SOJANODE", None)


@pytest.fixture
def basedir():
    return test_utilities.get_required_env_var("BASEDIR")


@pytest.fixture
def chain_id(is_ropsten_testnet):
    result = test_utilities.get_optional_env_var("DEPLOYMENT_NAME", "localnet")
    return result


@pytest.fixture
def ethereum_network_id(is_ropsten_testnet):
    result = test_utilities.get_optional_env_var("ETHEREUM_NETWORK_ID", None)
    if result:
        return result
    else:
        result = 3 if is_ropsten_testnet else 5777
        return result


@pytest.fixture
def ethereum_chain_id(ethereum_network_id):
    """For now, we always use the same network id and chain id"""
    return ethereum_network_id


@pytest.fixture
def ropsten_wait_time():
    return 30 * 60


@pytest.fixture
def is_ropsten_testnet(sojanoded_node):
    """if sojanode_clinode is set, we're talking to ropsten/sandpit"""
    return sojanoded_node


@pytest.fixture
def is_ganache(ethereum_network):
    """true if we're using ganache"""
    return not ethereum_network


# Deprecated: sojanoded accepts --gas-prices=0.5fury along with --gas-adjustment=1.5 instead of a fixed fee.
# Using those parameters is the best way to have the fees set robustly after the .42 upgrade.
# See https://github.com/Sojahub/sojanode/pull/1802#discussion_r697403408
@pytest.fixture
def sojahub_fees(sojahub_fees_int):
    """returns a string suitable for passing to sojanoded"""
    return f"{sojahub_fees_int}fury"


# Deprecated: sojanoded accepts --gas-prices=0.5fury along with --gas-adjustment=1.5 instead of a fixed fee.
# Using those parameters is the best way to have the fees set robustly after the .42 upgrade.
# See https://github.com/Sojahub/sojanode/pull/1802#discussion_r697403408
@pytest.fixture
def sojahub_fees_int():
    return 100000000000000000


@pytest.fixture
def operator_address(smart_contracts_dir):
    return test_utilities.get_optional_env_var("OPERATOR_ADDRESS",
                                               test_utilities.ganache_owner_account(smart_contracts_dir))


@pytest.fixture
def operator_private_key(ganache_keys_file, operator_address):
    result = test_utilities.get_optional_env_var(
        "OPERATOR_PRIVATE_KEY",
        test_utilities.ganache_private_key(ganache_keys_file, operator_address)
    )
    os.environ["OPERATOR_PRIVATE_KEY"] = result
    return result


@pytest.fixture
def ganache_keys_file(sojanode_base_dir):
    return test_utilities.get_optional_env_var(
        "GANACHE_KEYS_FILE",
        os.path.join(sojanode_base_dir, "test/integration/vagrant/data/ganachekeys.json")
    )


@pytest.fixture
def source_ethereum_address(is_ropsten_testnet, smart_contracts_dir):
    """
    Account with some starting eth that can be transferred out.

    Our test wallet can only use one address/privatekey combination,
    so if you set OPERATOR_ACCOUNT you have to set ETHEREUM_PRIVATE_KEY to the operator private key
    """
    addr = test_utilities.get_optional_env_var("ETHEREUM_ADDRESS", "")
    if addr:
        logging.debug("using ETHEREUM_ADDRESS provided for source_ethereum_address")
        return addr
    if is_ropsten_testnet:
        # Ropsten requires that you manually set the ETHEREUM_ADDRESS environment variable
        assert addr
    result = test_utilities.ganache_owner_account(smart_contracts_dir)
    logging.debug(
        f"Using source_ethereum_address {result} from ganache_owner_account.  (Set ETHEREUM_ADDRESS env var to set it manually)")
    assert result
    return result


@pytest.fixture(scope="function")
def ganache_timed_blocks(integration_dir):
    "restart ganache with timed blocks (keeps existing database)"
    logging.info("restart ganache with timed blocks (keeps existing database)")
    yield test_utilities.get_shell_output(f"{integration_dir}/ganache_start.sh 2")
    logging.info("restart ganache with instant mining (keeps existing database)")
    test_utilities.get_shell_output(f"{integration_dir}/ganache_start.sh")


@pytest.fixture(scope="function")
def ensure_relayer_restart(integration_dir, smart_contracts_dir):
    """restarts relayer after the test function completes.  Used by tests that need to stop the relayer."""
    yield None
    logging.info("restart ebrelayer after advancing wait blocks - avoids any interaction with replaying blocks")
    original_log_level = decrease_log_level(new_level=logging.WARNING)
    test_utilities.advance_n_ethereum_blocks(test_utilities.n_wait_blocks + 1, smart_contracts_dir)
    test_utilities.start_ebrelayer()
    force_log_level(original_log_level)


@pytest.fixture(scope="function")
def basic_transfer_request(
        smart_contracts_dir,
        bridgebank_address,
        bridgetoken_address,
        ethereum_network,
        sojanoded_node,
        chain_id,
        sojahub_fees,
        solidity_json_path,
        is_ganache,
):
    """
    Creates a EthereumToSojahubTransferRequest with all the generic fields filled in.
    """
    return test_utilities.EthereumToSojahubTransferRequest(
        smart_contracts_dir=smart_contracts_dir,
        ethereum_private_key_env_var="ETHEREUM_PRIVATE_KEY",
        bridgebank_address=bridgebank_address,
        bridgetoken_address=bridgetoken_address,
        ethereum_network=ethereum_network,
        sojanoded_node=sojanoded_node,
        manual_block_advance=is_ganache,
        chain_id=chain_id,
        sojahub_fees=sojahub_fees,
        solidity_json_path=solidity_json_path
    )


@pytest.fixture(scope="function")
def fury_source_integrationtest_env_credentials(
        sojanoded_homedir,
        validator_password,
        fury_source_key,
        is_ganache,
        fury_source
):
    """
    Creates a SojahubcliCredentials with all the fields filled in
    to transfer fury from an account that already has fury.
    """
    return test_utilities.SojahubcliCredentials(
        keyring_backend="test",
        keyring_passphrase=validator_password,
        from_key=fury_source
    )


@pytest.fixture(scope="function")
def fury_source_integrationtest_env_transfer_request(
        basic_transfer_request,
        fury_source
) -> test_utilities.EthereumToSojahubTransferRequest:
    """
    Creates a EthereumToSojahubTransferRequest with all the generic fields filled in
    for a transfer of fury from an account that already has fury.
    """
    result: test_utilities.EthereumToSojahubTransferRequest = copy.deepcopy(basic_transfer_request)
    result.sojahub_address = fury_source
    result.sojahub_symbol = "fury"
    return result


@pytest.fixture
def ethbridge_module_address():
    """The hardcoded address of the sojanode ethbridge module"""
    return "soja1l3dftf499u4gvdeuuzdl2pgv4f0xdtnuuwlzp8"


@pytest.fixture(scope="function")
def restore_default_rescue_location(
        ethbridge_module_address,
        sojahub_admin_account,
        sojahub_admin_account_credentials,
        basic_transfer_request
):
    """Restores the ethbridge module as the destination for ceth fees"""
    yield None
    test_utilities.update_ceth_receiver_account(
        receiver_account=ethbridge_module_address,
        admin_account=sojahub_admin_account,
        transfer_request=basic_transfer_request,
        credentials=sojahub_admin_account_credentials
    )


import sojatool_path
import sojatool.test_utils


@pytest.fixture(autouse=True)
def test_wrapper_fixture():
    sojatool.test_utils.pytest_test_wrapper_fixture()

@pytest.fixture(scope="function")
def ctx(request):
    # To pass the "snapshot_name" as a parameter with value "foo" from test, annotate the test function like this:
    # @pytest.mark.snapshot_name("foo")
    snapshot_name = request.node.get_closest_marker("snapshot_name")
    if snapshot_name is not None:
        snapshot_name = snapshot_name.args[0]
        logging.debug("Context setup: snapshot_name={}".format(repr(snapshot_name)))
    with sojatool.test_utils.get_test_env_ctx() as ctx:
        yield ctx
        logging.debug("Test context cleanup")
