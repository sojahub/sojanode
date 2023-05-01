#!/usr/bin/env bash

sojanoded tx clp set-lppd-params --path=./data/lpd_params.json \
	--from $ADMIN_KEY \
	--gas=500000 \
	--gas-prices=0.5fury \
	--chain-id $SOJAHUB_ID \
	--node $SOJANODE \
	--broadcast-mode block \
	--yes