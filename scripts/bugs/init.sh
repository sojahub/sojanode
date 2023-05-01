#!/usr/bin/env bash

### chain init script for development purposes only ###

make clean install
sojanoded init test --chain-id=localnet

echo "Generating deterministic account - soja"
echo "race draft rival universe maid cheese steel logic crowd fork comic easy truth drift tomorrow eye buddy head time cash swing swift midnight borrow" | sojanoded keys add soja --recover

echo "Generating deterministic account - akasha"
echo "hand inmate canvas head lunar naive increase recycle dog ecology inhale december wide bubble hockey dice worth gravity ketchup feed balance parent secret orchard" | sojanoded keys add akasha --recover

sojanoded add-genesis-account $(sojanoded keys show soja -a) 16205782692902021002506278400fury,500000000000000000000000catk,500000000000000000000000cbtk,500000000000000000000000ceth,990000000000000000000000000stake,500000000000000000000000cdash,500000000000000000000000clink,899999867990000000000000000000cacoin
sojanoded add-genesis-account $(sojanoded keys show akasha -a) 5000000000000003407464fury,500000000000000000000000catk,500000000000000000000000cbtk,500000000000000000000000ceth,990000000000000000000000000stake,500000000000000000000000cdash,500000000000000000000000clink,8999998679900000000000000000000cacoin

sojanoded add-genesis-clp-admin $(sojanoded keys show soja -a)
sojanoded add-genesis-clp-admin $(sojanoded keys show akasha -a)

sojanoded  add-genesis-validators $(sojanoded keys show soja -a --bech val)

sojanoded gentx soja 1000000000000000000000000stake --keyring-backend test

echo "Collecting genesis txs..."
sojanoded collect-gentxs

echo "Validating genesis file..."
sojanoded validate-genesis
