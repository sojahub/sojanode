#!/usr/bin/env bash

set -x

sojanoded tx clp cancel-unbond \
  --from $SOJA_ACT \
  --keyring-backend test \
  --symbol cusdt \
  --units 1000000000 \
  --fees 100000000000000000fury \
  --node ${SOJANODE_NODE} \
  --chain-id $SOJANODE_CHAIN_ID \
  --broadcast-mode block \
  -y