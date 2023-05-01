#!/usr/bin/env bash

set -x

sojanoded tx clp reward-params \
  --cancelPeriod 43200 \
  --lockPeriod 100800 \
  --from=$SOJA_ACT \
  --keyring-backend=test \
  --fees 100000000000000000fury \
  --gas 500000 \
  --node ${SOJANODE_NODE} \
  --chain-id=$SOJANODE_CHAIN_ID \
  --broadcast-mode=block \
  -y

# sojanoded tx clp reward-params \
#   --cancelPeriod 66825 \
#   --lockPeriod 124425 \
#   --from=$SOJA_ACT \
#   --keyring-backend=test \
#   --fees 100000000000000000fury \
#   --gas 500000 \
#   --node ${SOJANODE_NODE} \
#   --chain-id=$SOJANODE_CHAIN_ID \
#   --broadcast-mode=block \
#   -y

# sojanoded tx clp reward-params \
#   --cancelPeriod 66825 \
#   --lockPeriod 100800 \
#   --from=$SOJA_ACT \
#   --keyring-backend=test \
#   --fees 100000000000000000fury \
#   --gas 500000 \
#   --node ${SOJANODE_NODE} \
#   --chain-id=$SOJANODE_CHAIN_ID \
#   --broadcast-mode=block \
#   -y