#!/usr/bin/env bash

set -x

sojanoded q tokenregistry entries \
    --node ${SOJANODE_NODE} \
    --chain-id $SOJANODE_CHAIN_ID | jq