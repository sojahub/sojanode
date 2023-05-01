#!/usr/bin/env bash

set -x

sojanoded tx clp pmtp-params \
  --rGov=0.02 \
  --from=$SOJA_ACT \
  --keyring-backend=test \
  --fees 100000000000000000fury \
  --node ${SOJANODE_NODE} \
  --chain-id=$SOJANODE_CHAIN_ID \
  --broadcast-mode=block \
  -y