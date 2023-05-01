#!/usr/bin/env bash

sojanoded tx margin update-pools ./data/temp_pools.json \
	--closed-pools ./data/closed_pools.json \
  --from=$ADMIN_KEY \
	--gas=500000 \
	--gas-prices=0.5fury \
	--chain-id $SOJAHUB_ID \
	--node $SOJANODE \
	--broadcast-mode block \
	--yes

sojanoded tx margin whitelist soja1mwmrarhynjuau437d07p42803rntfxqjun3pfu \
  --from=$ADMIN_KEY \
	--gas=500000 \
	--gas-prices=0.5fury \
	--chain-id $SOJAHUB_ID \
	--node $SOJANODE \
	--broadcast-mode block \
	--yes