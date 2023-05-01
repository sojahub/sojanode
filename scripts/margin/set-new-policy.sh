#!/usr/bin/env bash

set -x

sojanoded tx clp pmtp-params \
  --pmtp_start=22811 \
  --pmtp_end=224410 \
  --epochLength=14400 \
  --rGov=0.05 \
  --from=$SOJA_ACT \
  --keyring-backend=test \
  --fees 100000000000000000fury \
  --node ${SOJANODE_NODE} \
  --chain-id=$SOJANODE_CHAIN_ID \
  --broadcast-mode=block \
  -y