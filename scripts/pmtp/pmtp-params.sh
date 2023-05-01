#!/usr/bin/env bash

set -x

sojanoded q clp pmtp-params \
  --node ${SOJANODE_NODE} \
  --chain-id $SOJANODE_CHAIN_ID