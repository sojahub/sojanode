#!/usr/bin/env bash

set -x

sojanoded tx margin close \
  --from $SOJA_ACT \
  --id 7 \
  --keyring-backend test \
  --fees 100000000000000000fury \
  --node ${SOJANODE_NODE} \
  --chain-id $SOJANODE_CHAIN_ID \
  --broadcast-mode block \
  -y