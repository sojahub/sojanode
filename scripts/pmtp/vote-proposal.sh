#!/usr/bin/env bash

set -x

sojanoded tx gov vote 1 yes \
    --from $SOJA_ACT \
    --keyring-backend test \
    --node ${SOJANODE_NODE} \
    --chain-id $SOJANODE_CHAIN_ID \
    --fees 100000000000000000fury \
    --broadcast-mode block \
    --trace \
    -y