#!/usr/bin/env bash

set -x

sojanoded q margin \
  positions-for-address $ADMIN_ADDRESS \
  --node ${SOJANODE_NODE} \
  --chain-id $SOJANODE_CHAIN_ID