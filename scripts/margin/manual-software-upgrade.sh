#!/usr/bin/env bash

set -x

sojanoded tx gov submit-proposal software-upgrade "${NEW_VERSION}" \
  --from ${SOJA_ACT} \
  --deposit "${DEPOSIT}" \
  --upgrade-height "${TARGET_BLOCK}" \
  --title "v${NEW_VERSION}" \
  --description "v${NEW_VERSION}" \
  --chain-id "${SOJANODE_CHAIN_ID}" \
  --node "${SOJANODE_NODE}" \
  --keyring-backend "test" \
  --fees 100000000000000000fury \
  --broadcast-mode=block \
  -y