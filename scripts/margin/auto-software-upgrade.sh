#!/usr/bin/env bash

set -x

sojanoded tx gov submit-proposal software-upgrade "${NEW_VERSION}" \
  --from ${SOJA_ACT} \
  --deposit "${DEPOSIT}" \
  --upgrade-height "${TARGET_BLOCK}" \
  --upgrade-info "{\"binaries\":{\"linux/amd64\":\"https://github.com/Sojahub/sojanode/releases/download/v${NEW_VERSION}/sojanoded-v${NEW_VERSION}-linux-amd64.zip?checksum=${CHECKSUM}\"}}" \
  --title "v${NEW_VERSION}" \
  --description "v${NEW_VERSION}" \
  --chain-id "${SOJANODE_CHAIN_ID}" \
  --node "${SOJANODE_NODE}" \
  --keyring-backend "test" \
  --fees 100000000000000000fury \
  --broadcast-mode=block \
  -y