#!/usr/bin/env bash

set -x

sojanoded tx gov submit-proposal \
    param-change proposal.json \
    --from $SOJA_ACT \
    --keyring-backend test \
    --node ${SOJANODE_NODE} \
    --chain-id $SOJANODE_CHAIN_ID \
    --fees 100000000000000000fury \
    --broadcast-mode block \
    -y