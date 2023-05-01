#!/usr/bin/env bash

set -x

sojanoded tx clp swap \
  --from $SOJA_ACT \
  --keyring-backend test \
  --sentSymbol fury \
  --receivedSymbol cusdt \
  --sentAmount 100000000000000000000000 \
  --minReceivingAmount 0 \
  --fees 100000000000000000fury \
  --node ${SOJANODE_NODE} \
  --chain-id $SOJANODE_CHAIN_ID \
  --broadcast-mode block \
  -y