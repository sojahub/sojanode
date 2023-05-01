#!/bin/sh

# sh ./deregister-all.sh testnet

. ./envs/$1.sh 

mkdir -p ./$SOJAHUB_ID
rm -f ./$SOJAHUB_ID/temp.json
rm -f ./$SOJAHUB_ID/temp2.json
rm -f ./$SOJAHUB_ID/tokenregistry.json

sojanoded q tokenregistry add-all ./$SOJAHUB_ID/registry.json | jq > $SOJAHUB_ID/temp.json
sojanoded q tokenregistry add ./$SOJAHUB_ID/temp.json ./$SOJAHUB_ID/cosmos.json | jq > $SOJAHUB_ID/tokenregistry.json
rm ./$SOJAHUB_ID/temp.json
sojanoded q tokenregistry add ./$SOJAHUB_ID/tokenregistry.json ./$SOJAHUB_ID/akash.json | jq > $SOJAHUB_ID/temp.json
rm ./$SOJAHUB_ID/tokenregistry.json
sojanoded q tokenregistry add ./$SOJAHUB_ID/temp.json ./$SOJAHUB_ID/sentinel.json | jq > $SOJAHUB_ID/tokenregistry.json
rm ./$SOJAHUB_ID/temp.json
sojanoded q tokenregistry add ./$SOJAHUB_ID/tokenregistry.json ./$SOJAHUB_ID/iris.json | jq > $SOJAHUB_ID/temp.json
rm ./$SOJAHUB_ID/tokenregistry.json
sojanoded q tokenregistry add ./$SOJAHUB_ID/temp.json ./$SOJAHUB_ID/persistence.json | jq > $SOJAHUB_ID/tokenregistry.json
rm ./$SOJAHUB_ID/temp.json
sojanoded q tokenregistry add ./$SOJAHUB_ID/tokenregistry.json ./$SOJAHUB_ID/crypto-org.json | jq > $SOJAHUB_ID/temp.json
rm ./$SOJAHUB_ID/tokenregistry.json
sojanoded q tokenregistry add ./$SOJAHUB_ID/temp.json ./$SOJAHUB_ID/regen.json | jq > $SOJAHUB_ID/tokenregistry.json
rm ./$SOJAHUB_ID/temp.json
sojanoded q tokenregistry add ./$SOJAHUB_ID/tokenregistry.json ./$SOJAHUB_ID/terra.json | jq > $SOJAHUB_ID/temp.json
rm ./$SOJAHUB_ID/tokenregistry.json
sojanoded q tokenregistry add ./$SOJAHUB_ID/temp.json ./$SOJAHUB_ID/osmosis.json | jq > $SOJAHUB_ID/tokenregistry.json
rm ./$SOJAHUB_ID/temp.json
sojanoded q tokenregistry add ./$SOJAHUB_ID/tokenregistry.json ./$SOJAHUB_ID/juno.json | jq > $SOJAHUB_ID/temp.json
rm ./$SOJAHUB_ID/tokenregistry.json
sojanoded q tokenregistry add ./$SOJAHUB_ID/temp.json ./$SOJAHUB_ID/ixo.json | jq > $SOJAHUB_ID/tokenregistry.json
rm ./$SOJAHUB_ID/temp.json
sojanoded q tokenregistry add ./$SOJAHUB_ID/tokenregistry.json ./$SOJAHUB_ID/emoney.json | jq > $SOJAHUB_ID/temp.json
rm ./$SOJAHUB_ID/tokenregistry.json
sojanoded q tokenregistry add ./$SOJAHUB_ID/temp.json ./$SOJAHUB_ID/likecoin.json | jq > $SOJAHUB_ID/tokenregistry.json
rm ./$SOJAHUB_ID/temp.json
sojanoded q tokenregistry add ./$SOJAHUB_ID/tokenregistry.json ./$SOJAHUB_ID/bitsong.json | jq > $SOJAHUB_ID/temp.json
rm ./$SOJAHUB_ID/tokenregistry.json
sojanoded q tokenregistry add ./$SOJAHUB_ID/temp.json ./$SOJAHUB_ID/band.json | jq > $SOJAHUB_ID/tokenregistry.json
rm ./$SOJAHUB_ID/temp.json
sojanoded q tokenregistry add ./$SOJAHUB_ID/tokenregistry.json ./$SOJAHUB_ID/emoney-eeur.json | jq > $SOJAHUB_ID/temp.json
rm ./$SOJAHUB_ID/tokenregistry.json
sojanoded q tokenregistry add ./$SOJAHUB_ID/temp.json ./$SOJAHUB_ID/terra-uusd.json | jq > $SOJAHUB_ID/tokenregistry.json
rm ./$SOJAHUB_ID/temp.json