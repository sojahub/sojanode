#!/usr/bin/env bash

set -x

sojanoded tx clp create-pool \
  --from $SOJA_ACT \
  --keyring-backend test \
  --symbol cusdt \
  --nativeAmount 1550459183129248235861408 \
  --externalAmount 174248776094 \
  --fees 100000000000000000fury \
  --node ${SOJANODE_NODE} \
  --chain-id $SOJANODE_CHAIN_ID \
  --broadcast-mode block \
  -y