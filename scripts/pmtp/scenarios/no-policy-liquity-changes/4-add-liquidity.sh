#!/usr/bin/env bash

set -x

sojanoded tx clp add-liquidity \
  --from $SOJA_ACT \
  --keyring-backend test \
  --symbol cusdt \
  --nativeAmount 100000000000000000000000 \
  --externalAmount 0 \
  --fees 100000000000000000fury \
  --node ${SOJANODE_NODE} \
  --chain-id $SOJANODE_CHAIN_ID \
  --broadcast-mode block \
  -y