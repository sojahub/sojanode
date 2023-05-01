#!/usr/bin/env bash

# Use sojanoded q account $(sojanoded keys show soja -a) to get seq
seq=1
sojanoded tx dispensation create Airdrop output.json --gas 90128 --from $(sojanoded keys show soja -a) --yes --broadcast-mode async --sequence $seq --account-number 3 --chain-id localnet
seq=$((seq+1))
sojanoded tx dispensation create ValidatorSubsidy output.json --gas 90128 --from $(sojanoded keys show soja -a) --yes --broadcast-mode async --sequence $seq --account-number 3 --chain-id localnet
seq=$((seq+1))
sojanoded tx dispensation create ValidatorSubsidy output.json --gas 90128 --from $(sojanoded keys show soja -a) --yes --broadcast-mode async --sequence $seq --account-number 3 --chain-id localnet