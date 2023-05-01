#!/usr/bin/env bash

set -x

sojanoded tx margin update-pools ./pools.json \
  --closed-pools ./closed-pools.json \
  --from=$SOJA_ACT \
  --keyring-backend=test \
  --fees 100000000000000000fury \
  --gas 500000 \
  --node ${SOJANODE_NODE} \
  --chain-id=$SOJANODE_CHAIN_ID \
  --broadcast-mode=block \
  -y