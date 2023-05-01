#!/usr/bin/env bash

pkill sojanoded
sleep 5
sojanoded export --height -1 > exported_state.json
sleep 1
sojanoded migrate v0.38 exported_state.json --chain-id new-chain > new-genesis.json  2>&1
sleep 1
sojanoded unsafe-reset-all
sleep 1
cp new-genesis.json ~/.sojanoded/config/genesis.json
sleep 2
sojanoded start