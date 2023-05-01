#!/usr/bin/env bash

### chain init script for development purposes only ###

make clean install
rm -rf ~/.sojanoded
sojanoded init test --chain-id=localnet -o

echo "Generating deterministic account - soja"
echo "race draft rival universe maid cheese steel logic crowd fork comic easy truth drift tomorrow eye buddy head time cash swing swift midnight borrow" | sojanoded keys add soja --recover --keyring-backend=test

echo "Generating deterministic account - akasha"
echo "hand inmate canvas head lunar naive increase recycle dog ecology inhale december wide bubble hockey dice worth gravity ketchup feed balance parent secret orchard" | sojanoded keys add akasha --recover --keyring-backend=test

echo "Generating deterministic account - alice"
echo "crunch enable gauge equip sadness venture volcano capable boil pole lounge because service level giggle decide south deposit bike antique consider olympic girl butter" | sojanoded keys add alice --recover --keyring-backend=test

sojanoded keys add mkey --multisig soja,akasha --multisig-threshold 2 --keyring-backend=test

sojanoded add-genesis-account $(sojanoded keys show soja -a --keyring-backend=test) 500000000000000000000000000000000fury,500000000000000000000000catk,500000000000000000000000cbtk,500000000000000000000000000000000ceth,990000000000000000000000000stake,500000000000000000000000cdash,500000000000000000000000clink,5000000000000cusdt,90000000000000000000ibc/96D7172B711F7F925DFC7579C6CCC3C80B762187215ABD082CDE99F81153DC80 --keyring-backend=test
sojanoded add-genesis-account $(sojanoded keys show akasha -a --keyring-backend=test) 500000000000000000000000fury,500000000000000000000000catk,500000000000000000000000cbtk,500000000000000000000000ceth,990000000000000000000000000stake,500000000000000000000000cdash,500000000000000000000000clink --keyring-backend=test
sojanoded add-genesis-account $(sojanoded keys show alice -a --keyring-backend=test) 500000000000000000000000fury,500000000000000000000000catk,500000000000000000000000cbtk,500000000000000000000000ceth,990000000000000000000000000stake,500000000000000000000000cdash,500000000000000000000000clink --keyring-backend=test

sojanoded add-genesis-clp-admin $(sojanoded keys show soja -a --keyring-backend=test) --keyring-backend=test
sojanoded add-genesis-clp-admin $(sojanoded keys show akasha -a --keyring-backend=test) --keyring-backend=test

sojanoded set-genesis-oracle-admin soja --keyring-backend=test
sojanoded add-genesis-validators $(sojanoded keys show soja -a --bech val --keyring-backend=test) --keyring-backend=test

sojanoded set-genesis-whitelister-admin soja --keyring-backend=test
sojanoded set-gen-denom-whitelist scripts/denoms.json

sojanoded gentx soja 1000000000000000000000000stake --moniker soja_val --chain-id=localnet --keyring-backend=test

echo "Collecting genesis txs..."
sojanoded collect-gentxs

echo "Validating genesis file..."
sojanoded validate-genesis
