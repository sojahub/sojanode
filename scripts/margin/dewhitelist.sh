#!/usr/bin/env bash

set -x

sojanoded tx margin dewhitelist did:fury:s1syavy2npfyt9tcncdtsdzf7kny9lh777p07psd \
  --from $SOJA_ACT \
  --keyring-backend test \
  --fees 100000000000000000fury \
  --node ${SOJANODE_NODE} \
  --chain-id $SOJANODE_CHAIN_ID \
  --broadcast-mode block \
  -y