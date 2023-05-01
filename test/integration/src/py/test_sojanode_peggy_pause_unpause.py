from typing import Tuple
import pytest

import sojatool_path
from sojatool import eth, test_utils, sojahub, cosmos
from sojatool.inflate_tokens import InflateTokens
from sojatool.common import *
from sojatool.test_utils import EnvCtx

def test_pause_unpause_no_error(ctx: EnvCtx):
    res = ctx.sojanode.pause_peggy_bridge(ctx.sojahub_ethbridge_admin_account)
    assert res[0]['code'] == 0
    res = ctx.sojanode.unpause_peggy_bridge(ctx.sojahub_ethbridge_admin_account)
    assert res[0]['code'] == 0

# We assert a tx is successful before pausing because we test the pause
# functionality by 1. An error response and 2. Balance unchanged within timeout.
# We want to make sure #2 is not a false positive due to lock function not
# working in the first place
def test_pause_lock_valid(ctx: EnvCtx):
    # Test a working flow:
    fund_amount_soja = 10 * test_utils.sojanode_funds_for_transfer_peggy1
    fund_amount_eth = 1 * eth.ETH

    test_soja_account = ctx.create_sojahub_addr(fund_amounts=[[fund_amount_soja, "fury"]])
    ctx.tx_bridge_bank_lock_eth(ctx.eth_faucet, test_soja_account, fund_amount_eth)
    ctx.eth.advance_blocks()
    # Setup is complete, test account has fury AND eth

    test_eth_destination_account = ctx.create_and_fund_eth_account()

    send_amount = 1
    balance_diff, erc_diff = send_test_account(ctx, test_soja_account, test_eth_destination_account, send_amount, erc20_token_addr=ctx.get_bridge_token_sc().address)
    # TODO: soja_tx_fee vs get from envctx vs more lenient assertion
    assert balance_diff.get(sojahub.FURY, 0) == (-1 * (send_amount + sojahub.soja_tx_fee_in_fury )), "Gas fee and sent amount should be deducted from sender soja acct"
    assert erc_diff == send_amount, "Eth destination should receive fury token"

    res = ctx.sojanode.pause_peggy_bridge(ctx.sojahub_ethbridge_admin_account)
    assert res[0]['code'] == 0

    balance_diff, erc_diff = send_test_account(ctx, test_soja_account, test_eth_destination_account, send_amount, erc20_token_addr=ctx.get_bridge_token_sc().address)
    assert balance_diff.get(sojahub.FURY, 0) == (-1 * sojahub.soja_tx_fee_in_fury), "Only gas fee should be deducted for attempted tx"
    assert erc_diff == 0, "Eth destination should not receive fury token"

    res = ctx.sojanode.unpause_peggy_bridge(ctx.sojahub_ethbridge_admin_account)
    assert res[0]['code'] == 0
    # # Submit lock
    # # Assert tx go through, balance updated correctly.
    send_amount = 15
    balance_diff, erc_diff = send_test_account(ctx, test_soja_account, test_eth_destination_account, send_amount, erc20_token_addr=ctx.get_bridge_token_sc().address)
    # TODO: soja_tx_fee vs get from envctx vs more lenient assertion
    assert balance_diff.get(sojahub.FURY, 0) == (-1 * (send_amount + sojahub.soja_tx_fee_in_fury )), "Gas fee and sent amount should be deducted from sender soja acct"
    assert erc_diff == send_amount, "Eth destination should receive fury token"

# Burn CETH
def test_pause_burn_valid(ctx: EnvCtx):
    fund_amount_soja = 10 * test_utils.sojanode_funds_for_transfer_peggy1
    fund_amount_eth = 1 * eth.ETH

    test_soja_account = ctx.create_sojahub_addr(fund_amounts=[[fund_amount_soja, "fury"]])
    soja_account_balance_before = ctx.get_sojahub_balance(test_soja_account)
    ctx.tx_bridge_bank_lock_eth(ctx.eth_faucet, test_soja_account, fund_amount_eth)
    ctx.eth.advance_blocks(100)
    # Setup is complete, test account has fury AND eth
    ctx.sojanode.wait_for_balance_change(test_soja_account, soja_account_balance_before)
    test_eth_destination_account = ctx.create_and_fund_eth_account()

    send_amount = 1
    balance_diff, erc_diff = send_test_account(ctx, test_soja_account, test_eth_destination_account, send_amount, denom=sojahub.CETH, erc20_token_addr=None)
    # TODO: soja_tx_fee vs get from envctx vs more lenient assertion
    gas_cost = 160000000000 * 393000 # Taken from peggy1
    assert balance_diff.get(sojahub.FURY, 0) == (-1 * sojahub.soja_tx_fee_in_fury), "Gas fee should be deducted from sender soja acct"
    assert balance_diff.get(sojahub.CETH, 0) == (-1 * (send_amount + gas_cost)), "Sent amount should be deducted from sender soja acct ceth balance"
    assert erc_diff == send_amount, "Eth destination should receive fury token"

    res = ctx.sojanode.pause_peggy_bridge(ctx.sojahub_ethbridge_admin_account)
    assert res[0]['code'] == 0

    send_amount = 1
    balance_diff, erc_diff = send_test_account(ctx, test_soja_account, test_eth_destination_account, send_amount, denom=sojahub.CETH, erc20_token_addr=None)
    assert balance_diff.get(sojahub.FURY, 0) == (-1 * sojahub.soja_tx_fee_in_fury), "Only gas fee should be deducted for attempted tx"
    assert balance_diff.get(sojahub.CETH, 0) == 0, "Eth amount should'nt be deducted, no tx to evm"
    assert erc_diff == 0, "Eth destination should not receive fury token"


    res = ctx.sojanode.unpause_peggy_bridge(ctx.sojahub_ethbridge_admin_account)
    assert res[0]['code'] == 0

    send_amount = 15
    balance_diff, erc_diff = send_test_account(ctx, test_soja_account, test_eth_destination_account, send_amount, denom=sojahub.CETH, erc20_token_addr=None)
    # TODO: soja_tx_fee vs get from envctx vs more lenient assertion
    gas_cost = 160000000000 * 393000 # Taken from peggy1
    assert balance_diff.get(sojahub.FURY, 0) == (-1 * sojahub.soja_tx_fee_in_fury), "Gas fee should be deducted from sender soja acct"
    assert balance_diff.get(sojahub.CETH, 0) == (-1 * (send_amount + gas_cost)), "Sent amount should be deducted from sender soja acct ceth balance"
    assert erc_diff == send_amount, "Eth destination should receive fury token"

# This is a temporary helper method. It will eventually be incorporated into sojatool
def send_test_account(ctx: EnvCtx, test_soja_account, test_eth_destination_account, send_amount, denom=sojahub.FURY, erc20_token_addr: str=None) -> Tuple[cosmos.Balance, int]:
    soja_balance_before = ctx.get_sojahub_balance(test_soja_account)
    if erc20_token_addr is not None:
        eth_balance_before = ctx.get_erc20_token_balance(erc20_token_addr, test_eth_destination_account)
    else:
        eth_balance_before = ctx.eth.get_eth_balance(test_eth_destination_account)

    ctx.sojanode_client.send_from_sojahub_to_ethereum(test_soja_account, test_eth_destination_account, send_amount, denom)

    soja_balance_after = ctx.sojanode.wait_for_balance_change(test_soja_account, soja_balance_before)
    try:
        eth_balance_after = ctx.wait_for_eth_balance_change(test_eth_destination_account, eth_balance_before, token_addr=erc20_token_addr, timeout=30)
    except Exception as e:
        # wait_for_eth_balance_change raises exception only if timedout, implying old_balance == new_balance
        eth_balance_after = eth_balance_before

    balance_diff = sojahub.balance_delta(soja_balance_before, soja_balance_after)
    return balance_diff, (eth_balance_after - eth_balance_before)