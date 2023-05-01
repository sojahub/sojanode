#!/usr/bin/env bash


clibuilder()
{
   echo ""
   echo "Usage: $0 -u UpgradeName -c CurrentBinary -n NewBinary"
   echo -e "\t-u Name of the upgrade [Must match a handler defined in setup-handlers.go in NewBinary]"
   echo -e "\t-c Branch name for old binary (Upgrade From)"
   echo -e "\t-n Branch name for new binary (Upgrade To)"
   exit 1 # Exit script after printing help
}

while getopts "u:c:n:" opt
do
   case "$opt" in
      u ) UpgradeName="$OPTARG" ;;
      c ) CurrentBinary="$OPTARG" ;;
      n ) NewBinary="$OPTARG" ;;
      ? ) clibuilder ;; # Print cliBuilder in case parameter is non-existent
   esac
done

if [ -z "$UpgradeName" ] || [ -z "$CurrentBinary" ] || [ -z "$NewBinary" ]
then
   echo "Some or all of the parameters are empty";
   clibuilder
fi


export DAEMON_HOME=$HOME/.sojanoded
export DAEMON_NAME=sojanoded
export DAEMON_ALLOW_DOWNLOAD_BINARIES=true

make clean
rm -rf ~/.sojanoded
rm -rf sojanode.log

rm -rf $GOPATH/bin/sojanoded
rm -rf $GOPATH/bin/old/sojanoded
rm -rf $GOPATH/bin/new/sojanoded

# Setup old binary and start chain
git checkout $CurrentBinary
make install
cp $GOPATH/bin/sojanoded $GOPATH/bin/old/
sojanoded init test --chain-id=localnet -o

echo "Generating deterministic account - soja"
echo "race draft rival universe maid cheese steel logic crowd fork comic easy truth drift tomorrow eye buddy head time cash swing swift midnight borrow" | sojanoded keys add soja --recover --keyring-backend=test

echo "Generating deterministic account - akasha"
echo "hand inmate canvas head lunar naive increase recycle dog ecology inhale december wide bubble hockey dice worth gravity ketchup feed balance parent secret orchard" | sojanoded keys add akasha --recover --keyring-backend=test


#sojanoded keys add mkey --multisig soja,akasha --multisig-threshold 2 --keyring-backend=test

sojanoded add-genesis-account $(sojanoded keys show soja -a --keyring-backend=test) 500000000000000000000000fury,500000000000000000000000catk,500000000000000000000000cbtk,500000000000000000000000ceth,990000000000000000000000000stake,500000000000000000000000cdash,500000000000000000000000clink,90000000000000000000ibc/96D7172B711F7F925DFC7579C6CCC3C80B762187215ABD082CDE99F81153DC80 --keyring-backend=test
sojanoded add-genesis-account $(sojanoded keys show akasha -a --keyring-backend=test) 500000000000000000000000fury,500000000000000000000000catk,500000000000000000000000cbtk,500000000000000000000000ceth,990000000000000000000000000stake,500000000000000000000000cdash,500000000000000000000000clink --keyring-backend=test

#sojanoded add-genesis-clp-admin $(sojanoded keys show soja -a --keyring-backend=test) --keyring-backend=test
#sojanoded add-genesis-clp-admin $(sojanoded keys show akasha -a --keyring-backend=test) --keyring-backend=test

#sojanoded set-genesis-whitelister-admin soja --keyring-backend=test

#sojanoded add-genesis-validators $(sojanoded keys show soja -a --bech val --keyring-backend=test) --keyring-backend=test

sojanoded gentx soja 1000000000000000000000000stake --chain-id=localnet --keyring-backend=test

echo "Collecting genesis txs..."
sojanoded collect-gentxs

echo "Validating genesis file..."
sojanoded validate-genesis


mkdir -p $DAEMON_HOME/cosmovisor/genesis/bin
mkdir -p $DAEMON_HOME/cosmovisor/upgrades/$UpgradeName/bin


# Setup new binary
git checkout $NewBinary
rm -rf $GOPATH/bin/sojanoded
make install
cp $GOPATH/bin/sojanoded $GOPATH/bin/new/


# Setup cosmovisor
cp $GOPATH/bin/old/sojanoded $DAEMON_HOME/cosmovisor/genesis/bin
cp $GOPATH/bin/new/sojanoded $DAEMON_HOME/cosmovisor/upgrades/$UpgradeName/bin/

chmod +x $DAEMON_HOME/cosmovisor/genesis/bin/sojanoded
chmod +x $DAEMON_HOME/cosmovisor/upgrades/$UpgradeName/bin/sojanoded

contents="$(jq '.app_state.gov.voting_params.voting_period = "10s"' $DAEMON_HOME/config/genesis.json)" && \
echo "${contents}" > $DAEMON_HOME/config/genesis.json

# Add state data here if required

cosmovisor start --home ~/.sojanoded/ --p2p.laddr 0.0.0.0:27655  --grpc.address 0.0.0.0:9096 --grpc-web.address 0.0.0.0:9093 --address tcp://0.0.0.0:27659 --rpc.laddr tcp://127.0.0.1:26657 >> sojanode.log 2>&1  &
#sleep 7
#sojanoded tx tokenregistry register-all /Users/tanmay/Documents/sojanode/scripts/ibc/tokenregistration/localnet/fury.json --from soja --keyring-backend=test --chain-id=localnet --yes
sleep 7
sojanoded tx gov submit-proposal software-upgrade $UpgradeName --from soja --deposit 100000000stake --upgrade-height 10 --title $UpgradeName --description $UpgradeName --keyring-backend test --chain-id localnet --yes
sleep 7
sojanoded tx gov vote 1 yes --from soja --keyring-backend test --chain-id localnet --yes
clear
sleep 7
sojanoded query gov proposal 1

tail -f sojanode.log

#yes Y | sojanoded tx gov submit-proposal software-upgrade 0.9.14 --from soja --deposit 100000000stake --upgrade-height 30 --title 0.9.14 --description 0.9.14 --keyring-backend test --chain-id localnet