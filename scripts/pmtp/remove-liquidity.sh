#!/usr/bin/env bash

set -x

sojanoded tx clp remove-liquidity \
  --from $SOJA_ACT \
  --keyring-backend test \
  --symbol cusdt \
  --asymmetry 0 \
  --wBasis 10 \
  --fees 100000000000000000fury \
  --node ${SOJANODE_NODE} \
  --chain-id $SOJANODE_CHAIN_ID \
  --broadcast-mode block \
  -y