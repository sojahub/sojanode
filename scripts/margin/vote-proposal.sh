#!/usr/bin/env bash

set -x

sojanoded tx gov vote 2 yes \
  --from ${SOJA_ACT} \
  --keyring-backend test \
  --chain-id="${SOJANODE_CHAIN_ID}" \
  --node="${SOJANODE_NODE}" \
  --fees=100000000000000000fury \
  --broadcast-mode=block \
  -y