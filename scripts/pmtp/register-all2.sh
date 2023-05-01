#!/usr/bin/env bash

set -x

sojanoded tx tokenregistry register denoms/cgala.json \
  --node ${SOJANODE_NODE} \
  --chain-id "${SOJANODE_CHAIN_ID}" \
  --from "${ADMIN_ADDRESS}" \
  --keyring-backend test \
  --gas 500000 \
  --gas-prices 0.5fury \
  -y \
  --broadcast-mode block

# sojanoded tx tokenregistry register denoms/fury.json \
#   --node ${SOJANODE_NODE} \
#   --chain-id "${SOJANODE_CHAIN_ID}" \
#   --from "${ADMIN_ADDRESS}" \
#   --keyring-backend test \
#   --gas 500000 \
#   --gas-prices 0.5fury \
#   -y \
#   --broadcast-mode block

# sojanoded tx tokenregistry register denoms/ceth.json \
#   --node ${SOJANODE_NODE} \
#   --chain-id "${SOJANODE_CHAIN_ID}" \
#   --from "${ADMIN_ADDRESS}" \
#   --keyring-backend test \
#   --gas 500000 \
#   --gas-prices 0.5fury \
#   -y \
#   --broadcast-mode block

# sojanoded tx tokenregistry register denoms/cusdc.json \
#   --node ${SOJANODE_NODE} \
#   --chain-id "${SOJANODE_CHAIN_ID}" \
#   --from "${ADMIN_ADDRESS}" \
#   --keyring-backend test \
#   --gas 500000 \
#   --gas-prices 0.5fury \
#   -y \
#   --broadcast-mode block

# sojanoded tx tokenregistry register denoms/cusdt.json \
#   --node ${SOJANODE_NODE} \
#   --chain-id "${SOJANODE_CHAIN_ID}" \
#   --from "${ADMIN_ADDRESS}" \
#   --keyring-backend test \
#   --gas 500000 \
#   --gas-prices 0.5fury \
#   -y \
#   --broadcast-mode block

# sojanoded tx tokenregistry register denoms/uatom.json \
#   --node ${SOJANODE_NODE} \
#   --chain-id "${SOJANODE_CHAIN_ID}" \
#   --from "${ADMIN_ADDRESS}" \
#   --keyring-backend test \
#   --gas 500000 \
#   --gas-prices 0.5fury \
#   -y \
#   --broadcast-mode block

# sojanoded tx tokenregistry register denoms/ujuno.json \
#   --node ${SOJANODE_NODE} \
#   --chain-id "${SOJANODE_CHAIN_ID}" \
#   --from "${ADMIN_ADDRESS}" \
#   --keyring-backend test \
#   --gas 500000 \
#   --gas-prices 0.5fury \
#   -y \
#   --broadcast-mode block

# sojanoded tx tokenregistry register denoms/uluna.json \
#   --node ${SOJANODE_NODE} \
#   --chain-id "${SOJANODE_CHAIN_ID}" \
#   --from "${ADMIN_ADDRESS}" \
#   --keyring-backend test \
#   --gas 500000 \
#   --gas-prices 0.5fury \
#   -y \
#   --broadcast-mode block