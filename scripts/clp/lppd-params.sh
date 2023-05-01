#!/bin/sh

set -x

sojanoded tx clp set-lppd-params \
  --path lppd-params.json \
  --from $SOJA_ACT \
  --keyring-backend test \
  --fees 100000000000000000fury \
  --node ${SOJANODE_NODE} \
  --chain-id $SOJANODE_CHAIN_ID \
  --broadcast-mode block \
  -y