#!/usr/bin/env bash

. ../credentials.sh

rm -rf ~/.sojanoded

sojanoded init test --chain-id=sojahub-local
cp ./app.toml ~/.sojanoded/config

echo "Generating deterministic account - ${SHADOWFIEND_NAME}"
echo "${SHADOWFIEND_MNEMONIC}" | sojanoded keys add ${SHADOWFIEND_NAME}  --keyring-backend=test --recover

echo "Generating deterministic account - ${AKASHA_NAME}"
echo "${AKASHA_MNEMONIC}" | sojanoded keys add ${AKASHA_NAME}  --keyring-backend=test --recover

echo "Generating deterministic account - ${JUNIPER_NAME}"
echo "${JUNIPER_MNEMONIC}" | sojanoded keys add ${JUNIPER_NAME} --keyring-backend=test --recover

sojanoded add-genesis-account $(sojanoded keys show ${SHADOWFIEND_NAME} -a --keyring-backend=test) 100000000000000000000000000000fury,100000000000000000000000000000catk,100000000000000000000000000000cbtk,100000000000000000000000000000ceth,100000000000000000000000000000cusdc,100000000000000000000000000000clink,100000000000000000000000000stake
sojanoded add-genesis-account $(sojanoded keys show ${AKASHA_NAME} -a --keyring-backend=test) 100000000000000000000000000000fury,100000000000000000000000000000catk,100000000000000000000000000000cbtk,100000000000000000000000000000ceth,100000000000000000000000000000cusdc,100000000000000000000000000000clink,100000000000000000000000000stake
sojanoded add-genesis-account $(sojanoded keys show ${JUNIPER_NAME} -a --keyring-backend=test) 10000000000000000000000fury,10000000000000000000000cusdc,100000000000000000000clink,100000000000000000000ceth

sojanoded add-genesis-validators $(sojanoded keys show ${SHADOWFIEND_NAME} -a --bech val --keyring-backend=test)

sojanoded gentx ${SHADOWFIEND_NAME} 1000000000000000000000000stake --chain-id=sojahub-local --keyring-backend test

echo "Collecting genesis txs..."
sojanoded collect-gentxs

echo "Validating genesis file..."
sojanoded validate-genesis

echo "Starting test chain"

./start.sh
