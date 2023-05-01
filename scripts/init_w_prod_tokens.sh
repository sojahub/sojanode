#!/usr/bin/env bash

### chain init script for development purposes only ###

make clean install
rm -rf ~/.sojanoded
sojanoded init test --chain-id=localnet -o

echo "Generating deterministic account - soja"
echo "race draft rival universe maid cheese steel logic crowd fork comic easy truth drift tomorrow eye buddy head time cash swing swift midnight borrow" | sojanoded keys add soja --recover --keyring-backend=test

echo "Generating deterministic account - akasha"
echo "hand inmate canvas head lunar naive increase recycle dog ecology inhale december wide bubble hockey dice worth gravity ketchup feed balance parent secret orchard" | sojanoded keys add akasha --recover --keyring-backend=test


sojanoded keys add mkey --multisig soja,akasha --multisig-threshold 2 --keyring-backend=test

sojanoded add-genesis-account $(sojanoded keys show soja -a --keyring-backend=test) "999000000000000000000000000000000fury,999000000000000000000000000000000stake,999000000000000000000000000000000ceth,999000000000000000000000000000000cusdc,999000000000000000000000000000000cusdt,999000000000000000000000000000000ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2,999000000000000000000000000000000ibc/F279AB967042CAC10BFF70FAECB179DCE37AAAE4CD4C1BC4565C2BBC383BC0FA,999000000000000000000000000000000ibc/F141935FF02B74BDC6B8A0BD6FE86A23EE25D10E89AA0CD9158B3D92B63FDF4D" --keyring-backend=test
sojanoded add-genesis-account $(sojanoded keys show akasha -a --keyring-backend=test) 500000000000000000000000fury,500000000000000000000000catk,500000000000000000000000cbtk,500000000000000000000000ceth,990000000000000000000000000stake,500000000000000000000000cdash,500000000000000000000000clink --keyring-backend=test

sojanoded add-genesis-clp-admin $(sojanoded keys show soja -a --keyring-backend=test) --keyring-backend=test
sojanoded add-genesis-clp-admin $(sojanoded keys show akasha -a --keyring-backend=test) --keyring-backend=test

sojanoded set-genesis-oracle-admin soja --keyring-backend=test
sojanoded add-genesis-validators $(sojanoded keys show soja -a --bech val --keyring-backend=test) --keyring-backend=test

# FIXME: commented as it overrides admin accounts list in genesis set by default
# sojanoded set-genesis-whitelister-admin soja --keyring-backend=test
# sojanoded set-gen-denom-whitelist scripts/denoms.json

sojanoded gentx soja 1000000000000000000000000stake --chain-id=localnet --keyring-backend=test

echo "Collecting genesis txs..."
sojanoded collect-gentxs

echo "Validating genesis file..."
sojanoded validate-genesis
