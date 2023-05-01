#!/bin/zsh

proposal_count=1
sojanoded q gov proposals --node tcp://127.0.0.1:27665 >/dev/null 2>&1
if [ $? -eq 0 ]
then
  proposal_count=2
fi

echo "vote for $proposal_count on localnet-2"
sojanoded tx gov submit-proposal update-client 07-tendermint-0 07-tendermint-2 --from soja --keyring-backend test --home ~/.sojanode-2  --node tcp://127.0.0.1:27666 --title "vote for $proposal_count" --description "vote for $proposal_count" --chain-id localnet-2  --deposit 100000000stake --broadcast-mode block --yes 
echo "proposal made"
sojanoded tx gov vote $proposal_count yes --chain-id localnet-2 --from soja --keyring-backend test --home ~/.sojanode-2 --node tcp://127.0.0.1:27666 --yes --broadcast-mode block 
echo "sleeping"
sleep 30
sojanoded tx gov vote $proposal_count yes --chain-id localnet-2 --from akasha --keyring-backend test --home ~/.sojanode-2 --node tcp://127.0.0.1:27666 --yes --broadcast-mode block 
echo "done"


#echo "vote for $proposal_count on localnet-3"
sojanoded tx gov submit-proposal update-client 07-tendermint-1 07-tendermint-2 --from soja --keyring-backend test --home ~/.sojanode-3  --node tcp://127.0.0.1:27667 --title "vote for $proposal_count" --description "vote for $proposal_count" --chain-id localnet-3  --deposit 100000000stake --broadcast-mode block --yes 
echo "proposal made"
sojanoded tx gov vote $proposal_count yes --chain-id localnet-3 --from soja --keyring-backend test --home ~/.sojanode-3 --node tcp://127.0.0.1:27667 --yes --broadcast-mode block 
echo "sleeping"
sleep 30
sojanoded tx gov vote $proposal_count yes --chain-id localnet-3 --from akasha --keyring-backend test --home ~/.sojanode-3 --node tcp://127.0.0.1:27667 --yes --broadcast-mode block 
echo "done"
