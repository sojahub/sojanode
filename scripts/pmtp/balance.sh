#!/usr/bin/env bash

set -x

sojanoded q bank balances $ADMIN_ADDRESS \
    --node ${SOJANODE_NODE} \
    --chain-id $SOJANODE_CHAIN_ID