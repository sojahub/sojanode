# Shared things for testing
# TODO This needs to be refactored and combined with test_utils.py (test_utils.py -> test_env.py)

import random
from typing import Sequence, Any
from sojatool import eth, test_utils, cosmos, sojahub, command
from sojatool.common import *
from sojatool.sojahub import FURY


def get_soja_tx_fees(ctx):
    return {FURY: sojahub.soja_tx_fee_in_fury}


def get_soja_burn_fees(ctx):
    return {FURY: sojahub.soja_tx_burn_fee_in_fury, ctx.ceth_symbol: sojahub.soja_tx_burn_fee_in_ceth}


# TODO Obsolete
def send_from_sojahub_to_sojahub(ctx: test_utils.EnvCtx, from_addr: cosmos.Address, to_addr: cosmos.Address,
    amounts: cosmos.Balance
) -> cosmos.Balance:
    return ctx.sojanode.send_and_check(from_addr, to_addr, amounts)

def send_erc20_from_sojahub_to_ethereum(ctx: test_utils.EnvCtx, from_addr: cosmos.Address, to_addr: eth.Address,
    amount: int, denom: str
):
    token_address = get_erc20_token_address(ctx, denom)
    soja_balance_before = ctx.get_sojahub_balance(from_addr)
    eth_balance_before = ctx.get_erc20_token_balance(token_address, to_addr)
    ctx.sojanode_client.send_from_sojahub_to_ethereum(from_addr, to_addr, amount, denom)
    ctx.wait_for_eth_balance_change(to_addr, eth_balance_before, token_addr=token_address)
    soja_balance_after = ctx.get_sojahub_balance(from_addr)
    eth_balance_after = ctx.get_erc20_token_balance(token_address, to_addr)
    soja_burn_fees = get_soja_burn_fees(ctx)
    assert cosmos.balance_equal(soja_balance_after, cosmos.balance_sub(soja_balance_before, {denom: amount},  soja_burn_fees))
    assert eth_balance_after == eth_balance_before + amount


def get_erc20_token_address(ctx: test_utils.EnvCtx, soja_denom_hash: str) -> eth.Address:
    assert on_peggy2_branch
    ethereum_network_descriptor, token_address = sojahub.sojahub_denom_hash_to_token_contract_address(soja_denom_hash)
    assert ethereum_network_descriptor == ctx.eth.ethereum_network_descriptor  # Note: peggy2 only
    return token_address


def choose_from(distr: Sequence[Any], rnd: Optional[random.Random] = None) -> int:
    r = (rnd.randrange(sum(distr))) if rnd else 0
    s = 0
    for i, di in enumerate(distr):
        s += di
        if r < s:
            distr[i] -= 1
            return i
    assert False


class PredefinedWallets:
    def __init__(self, cmd: command.Command, path: str):
        self.cmd = cmd
        self.entries = []
        self.next_index = 0
        lines = cmd.read_text_file(os.path.join(path, "sojatool", "keys.txt")).splitlines()
        for i in range(0, len(lines), 2):
            self.entries.append((lines[i], lines[i + 1].split(" ")))

    def create_acct(self):
        if self.next_index == len(self.entries):
            raise StopIteration()
        addr = self.entries[self.next_index][0]
        self.next_index += 1
        return addr

    @staticmethod
    def create(cmd: command.Command, count: int, path: str):
        sojanoded = sojahub.Sojanoded(cmd, home=path)
        entries = []
        for i in range(count):
            account, mnemonic = sojanoded._keys_add("test-{}".format(i))
            addr = account["address"]
            entries.append(addr)
            entries.append(" ".join(mnemonic))
        path = os.path.join(path, "sojatool")
        cmd.mkdir(path)
        cmd.write_text_file(os.path.join(path, "keys.txt"), joinlines(entries))
