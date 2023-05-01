#!/usr/bin/env bash

set -x

sojanoded tx margin whitelist $(sojanoded keys show tester1 --keyring-backend=test -a) \
  --from $ADMIN_KEY \
  --keyring-backend test \
  --fees 100000000000000000fury \
  --node ${SOJANODE_NODE} \
  --chain-id $SOJANODE_CHAIN_ID \
  --broadcast-mode block \
  -y
sojanoded tx margin whitelist $(sojanoded keys show tester2 --keyring-backend=test -a) \
  --from $ADMIN_KEY \
  --keyring-backend test \
  --fees 100000000000000000fury \
  --node ${SOJANODE_NODE} \
  --chain-id $SOJANODE_CHAIN_ID \
  --broadcast-mode block \
  -y
sojanoded tx margin whitelist $(sojanoded keys show tester3 --keyring-backend=test -a) \
  --from $ADMIN_KEY \
  --keyring-backend test \
  --fees 100000000000000000fury \
  --node ${SOJANODE_NODE} \
  --chain-id $SOJANODE_CHAIN_ID \
  --broadcast-mode block \
  -y
sojanoded tx margin whitelist $(sojanoded keys show tester4 --keyring-backend=test -a) \
  --from $ADMIN_KEY \
  --keyring-backend test \
  --fees 100000000000000000fury \
  --node ${SOJANODE_NODE} \
  --chain-id $SOJANODE_CHAIN_ID \
  --broadcast-mode block \
  -y
sojanoded tx margin whitelist $(sojanoded keys show tester5 --keyring-backend=test -a) \
  --from $ADMIN_KEY \
  --keyring-backend test \
  --fees 100000000000000000fury \
  --node ${SOJANODE_NODE} \
  --chain-id $SOJANODE_CHAIN_ID \
  --broadcast-mode block \
  -y
sojanoded tx margin whitelist $(sojanoded keys show tester6 --keyring-backend=test -a) \
  --from $ADMIN_KEY \
  --keyring-backend test \
  --fees 100000000000000000fury \
  --node ${SOJANODE_NODE} \
  --chain-id $SOJANODE_CHAIN_ID \
  --broadcast-mode block \
  -y
sojanoded tx margin whitelist $(sojanoded keys show tester7 --keyring-backend=test -a) \
  --from $ADMIN_KEY \
  --keyring-backend test \
  --fees 100000000000000000fury \
  --node ${SOJANODE_NODE} \
  --chain-id $SOJANODE_CHAIN_ID \
  --broadcast-mode block \
  -y
sojanoded tx margin whitelist $(sojanoded keys show tester8 --keyring-backend=test -a) \
  --from $ADMIN_KEY \
  --keyring-backend test \
  --fees 100000000000000000fury \
  --node ${SOJANODE_NODE} \
  --chain-id $SOJANODE_CHAIN_ID \
  --broadcast-mode block \
  -y
sojanoded tx margin whitelist $(sojanoded keys show tester9 --keyring-backend=test -a) \
  --from $ADMIN_KEY \
  --keyring-backend test \
  --fees 100000000000000000fury \
  --node ${SOJANODE_NODE} \
  --chain-id $SOJANODE_CHAIN_ID \
  --broadcast-mode block \
  -y
sojanoded tx margin whitelist $(sojanoded keys show tester10 --keyring-backend=test -a) \
  --from $ADMIN_KEY \
  --keyring-backend test \
  --fees 100000000000000000fury \
  --node ${SOJANODE_NODE} \
  --chain-id $SOJANODE_CHAIN_ID \
  --broadcast-mode block \
  -y
sojanoded tx margin whitelist $(sojanoded keys show tester11 --keyring-backend=test -a) \
  --from $ADMIN_KEY \
  --keyring-backend test \
  --fees 100000000000000000fury \
  --node ${SOJANODE_NODE} \
  --chain-id $SOJANODE_CHAIN_ID \
  --broadcast-mode block \
  -y
