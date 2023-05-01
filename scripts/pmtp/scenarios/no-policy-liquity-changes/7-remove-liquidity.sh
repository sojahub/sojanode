#!/usr/bin/env bash

set -x

sojanoded tx clp remove-liquidity \
  --from $SOJA_ACT \
  --keyring-backend test \
  --symbol cusdt \
  --asymmetry 10000 \
  --wBasis 295 \
  --fees 100000000000000000fury \
  --node ${SOJANODE_NODE} \
  --chain-id $SOJANODE_CHAIN_ID \
  --broadcast-mode block \
  -y