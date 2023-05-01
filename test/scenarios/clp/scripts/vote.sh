#!/bin/sh

# Vote yes to accept the proposal
sojanoded tx gov vote 1 yes \
--from soja --keyring-backend test \
--fees 100000fury \
--chain-id  localnet \
--broadcast-mode block \
-y