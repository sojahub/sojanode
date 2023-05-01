#!/usr/bin/env bash

set -x

sojanoded tx clp pmtp-rates \
  --blockRate=0.00 \
  --runningRate=0.00 \
  --from=$SOJA_ACT \
  --keyring-backend=test \
  --fees 100000000000000000fury \
  --node ${SOJANODE_NODE} \
  --chain-id=$SOJANODE_CHAIN_ID \
  --broadcast-mode=block \
  -y