#!/usr/bin/env bash

set -x

ACCOUNT_NUMBER=$(sojanoded q auth account $ADMIN_ADDRESS \
    --node ${SOJANODE_NODE} \
    --chain-id $SOJANODE_CHAIN_ID \
    --output json \
    | jq -r ".account_number")
SEQUENCE=$(sojanoded q auth account $ADMIN_ADDRESS \
  --node ${SOJANODE_NODE} \
  --chain-id $SOJANODE_CHAIN_ID \
  --output json \
  | jq -r ".sequence")
for i in {0..12244}; do
  echo "tx ${i}"
  sojanoded tx clp add_liquidity \
    --from=$SOJA_ACT \
    --keyring-backend=test \
    --externalAmount=${EXTERNAL_AMOUNT} \
    --nativeAmount=${NATIVE_AMOUNT} \
    --symbol=${SYMBOL} \
    --fees=100000000000000000fury \
    --gas=500000 \
    --node=${SOJANODE_NODE} \
    --chain-id=${SOJANODE_CHAIN_ID} \
    --broadcast-mode=async \
    --account-number=${ACCOUNT_NUMBER} \
    --sequence=$(($SEQUENCE + $i)) \
    -y
  done