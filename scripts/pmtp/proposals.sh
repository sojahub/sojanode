#!/usr/bin/env bash

set -x

sojanoded q gov proposals \
    --node ${SOJANODE_NODE} \
    --chain-id $SOJANODE_CHAIN_ID