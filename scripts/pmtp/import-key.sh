#!/usr/bin/env bash

set -x

echo ${ADMIN_MNEMONIC} | sojanoded keys add ${SOJA_ACT} --recover --keyring-backend=test