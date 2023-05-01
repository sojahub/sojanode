#!/usr/bin/env bash

sojanoded tx clp reward-period --path=./data/atom_rewards_fix.json \
	--from $ADMIN_KEY \
	--gas=500000 \
	--gas-prices=0.5fury \
	--chain-id $SOJAHUB_ID \
	--node $SOJANODE \
	--broadcast-mode block \
	--yes