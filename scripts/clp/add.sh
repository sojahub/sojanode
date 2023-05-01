#!/bin/sh

set -x

sojanoded tx clp add-liquidity \
  --externalAmount 488436982990 \
  --nativeAmount 96176925423929435353999282 \
  --symbol ceth \
  --from $SOJA_ACT \
  --keyring-backend test \
  --fees 100000000000000000fury \
  --node ${SOJANODE_NODE} \
  --chain-id $SOJANODE_CHAIN_ID \
  --broadcast-mode block \
  -y