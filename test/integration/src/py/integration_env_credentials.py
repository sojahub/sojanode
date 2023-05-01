import burn_lock_functions
from burn_lock_functions import SojahubcliCredentials
from test_utilities import get_required_env_var, get_shell_output


def sojahub_cli_credentials_for_test(key: str) -> SojahubcliCredentials:
    """Returns SojahubcliCredentials for the test keyring with from_key set to key"""
    return SojahubcliCredentials(
        keyring_passphrase="",
        keyring_backend="test",
        from_key=key,
        sojanoded_homedir=f"""{get_required_env_var("HOME")}/.sojanoded"""
    )


def create_new_sojaaddr_and_credentials() -> (str, SojahubcliCredentials):
    new_account_key = get_shell_output("uuidgen")
    credentials = sojahub_cli_credentials_for_test(new_account_key)
    new_addr = burn_lock_functions.create_new_sojaaddr(credentials=credentials, keyname=new_account_key)
    credentials.from_key = new_addr["name"]
    return new_addr["address"], credentials,
