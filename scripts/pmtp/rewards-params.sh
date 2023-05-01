#!/usr/bin/env bash

set -x

sojanoded q clp reward-params \
  --node ${SOJANODE_NODE} \
  --chain-id $SOJANODE_CHAIN_ID