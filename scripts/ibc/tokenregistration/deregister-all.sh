#!/bin/sh

# sh ./deregister-all.sh testnet

. ./envs/$1.sh 

TOKEN_REGISTRY_ADMIN_ADDRESS="did:fury:s1tpypxpppcf5lea47vcvgy09675nllmcucxydvu"

sojanoded tx tokenregistry deregister-all ./$SOJAHUB_ID/tokenregistry.json \
  --node $SOJA_NODE \
  --chain-id $SOJAHUB_ID \
  --from $TOKEN_REGISTRY_ADMIN_ADDRESS \
  --keyring-backend $KEYRING_BACKEND \
  --gas=500000 \
  --gas-prices=0.5fury \
  -y