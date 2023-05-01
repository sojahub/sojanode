#!/usr/bin/env bash

set -x

sojanoded tx clp set-symmetry-threshold \
  --threshold=0.000000005 \
  --from=$SOJA_ACT \
  --keyring-backend=test \
  --fees=100000000000000000fury \
  --gas=500000 \
  --node=${SOJANODE_NODE} \
  --chain-id=$SOJANODE_CHAIN_ID \
  --broadcast-mode=block \
  -y