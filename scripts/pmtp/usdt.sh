#!/usr/bin/env bash

set -x

sojanoded q clp pool cusdt \
  --node ${SOJANODE_NODE} \
  --chain-id $SOJANODE_CHAIN_ID