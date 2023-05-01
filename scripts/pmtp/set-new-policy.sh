#!/usr/bin/env bash

set -x

sojanoded tx clp pmtp-params \
  --pmtp_start=31 \
  --pmtp_end=1030 \
  --epochLength=100 \
  --rGov=0.10 \
  --from=$SOJA_ACT \
  --keyring-backend=test \
  --fees 100000000000000000fury \
  --node ${SOJANODE_NODE} \
  --chain-id=$SOJANODE_CHAIN_ID \
  --broadcast-mode=block \
  -y