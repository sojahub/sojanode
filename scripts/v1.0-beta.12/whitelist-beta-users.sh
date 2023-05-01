#!/usr/bin/env bash

source ./data/margin-beta-users.sh

for addr in $users
do
  sojanoded tx margin whitelist $addr \
    --from=$ADMIN_KEY \
  	--gas=500000 \
  	--gas-prices=0.5fury \
  	--chain-id $SOJAHUB_ID \
  	--node $SOJANODE \
  	--broadcast-mode block \
  	--yes
done