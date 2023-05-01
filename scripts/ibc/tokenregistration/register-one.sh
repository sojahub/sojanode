#!/bin/sh
. ./envs/$1.sh 

# sh ./register-one.sh testnet ixo


TOKEN_REGISTRY_ADMIN_ADDRESS="soja1tpypxpppcf5lea47vcvgy09675nllmcucxydvu"

sojanoded tx tokenregistry register ./$SOJAHUB_ID/$2.json \
  --node $SOJA_NODE \
  --chain-id $SOJAHUB_ID \
  --from $TOKEN_REGISTRY_ADMIN_ADDRESS \
  --keyring-backend $KEYRING_BACKEND \
  --gas=500000 \
  --gas-prices=0.5fury \
  -y