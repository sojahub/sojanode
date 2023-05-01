#!/bin/sh

# Create fury/ceth; 
sojanoded tx clp create-pool \
--symbol ceth \
--nativeAmount 2000000000000000000 \
--externalAmount 2000000000000000000 \
--from soja --keyring-backend test \
--fees 100000000000000000fury \
--chain-id localnet \
--broadcast-mode block \
-y