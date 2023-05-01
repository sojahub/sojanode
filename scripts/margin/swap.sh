#!/usr/bin/env bash

set -x

sojanoded tx clp swap \
  --from=$SOJA_ACT \
  --keyring-backend=test \
  --sentSymbol=cusdc \
  --receivedSymbol=fury \
  --sentAmount=1000000000000 \
  --minReceivingAmount=0 \
  --fees=100000000000000000fury \
  --gas=500000 \
  --node=${SOJANODE_NODE} \
  --chain-id=${SOJANODE_CHAIN_ID} \
  --broadcast-mode=block \
  -y