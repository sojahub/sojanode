#!/usr/bin/env bash

set -x

sojanoded q clp all-lp \
  --node ${SOJANODE_NODE} \
  --chain-id $SOJANODE_CHAIN_ID