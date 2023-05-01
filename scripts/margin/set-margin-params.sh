#!/usr/bin/env bash

set -x

sojanoded tx margin update-params \
  --health-gain-factor=0.0000000001 \
  --interest-rate-decrease=0.000000001 \
  --interest-rate-increase=0.000000001 \
  --interest-rate-max=3.0 \
  --interest-rate-min=0.000000000005 \
  --leverage-max=2 \
  --epoch-length=1 \
  --removal-queue-threshold=0.1 \
  --max-open-positions=10000 \
  --force-close-fund-percentage=0.1 \
  --force-close-fund-address=did:fury:s1syavy2npfyt9tcncdtsdzf7kny9lh777p07psd \
  --incremental-interest-payment-enabled=true \
  --incremental-interest-payment-fund-percentage=0.1 \
  --incremental-interest-payment-fund-address=did:fury:s1syavy2npfyt9tcncdtsdzf7kny9lh777p07psd \
  --pool-open-threshold=0.0000000000001 \
  --sq-modifier=10000000000000000000000000 \
  --safety-factor=1.05 \
  --whitelisting-enabled=true \
  --from=$SOJA_ACT \
  --keyring-backend=test \
  --fees 100000000000000000fury \
  --gas 500000 \
  --node ${SOJANODE_NODE} \
  --chain-id=$SOJANODE_CHAIN_ID \
  --broadcast-mode=block \
  -y