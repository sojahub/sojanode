#!/usr/bin/env bash

set -x

sojanoded tx bank send \
    $SOJA_ACT \
    soja144w8cpva2xkly74xrms8djg69y3mljzplx3fjt \
    9299999999750930000fury \
    --keyring-backend test \
    --node ${SOJANODE_NODE} \
    --chain-id $SOJANODE_CHAIN_ID \
    --fees 100000000000000000fury \
    --broadcast-mode block \
    -y