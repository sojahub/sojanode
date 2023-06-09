required="${BASEDIR:?"Must be set to where you check out sojanode"}"

cd $BASEDIR/test/integration

required="${DEPLOYMENT_NAME:?"Must be set to a deployment name like sandpit"}"
required="${FURY_SOURCE:?"Must be set to a soja address that contains fury, and the key must be in your keyring"}"
required="${ETHEREUM_NETWORK:?"Must be set to an etherereum endpoint.  Can be ropsten or a url."}"
required="${ETHEREUM_PRIVATE_KEY:?"Must be set to the private key of the address specified in ETHEREUM_ADDRESS or OPERATOR_ADDRESS"}"
required="${SOJANODE:?"Must be set to sojanode endpoint."}"
required="${INFURA_PROJECT_ID:?"Must be set."}"

if [ -z "$ETHEREUM_ADDRESS$OPERATOR_ADDRESS" ]; then echo must set one of ETHEREUM_ADDRESS or OPERATOR_ADDRESS; exit 1; fi

export SMART_CONTRACTS_DIR=$BASEDIR/smart-contracts
export SOLIDITY_JSON_PATH=$BASEDIR/smart-contracts/deployments/$DEPLOYMENT_NAME
export SMART_CONTRACT_ARTIFACT_DIR=$SOLIDITY_JSON_PATH

export BRIDGE_REGISTRY_ADDRESS=$(cat $SOLIDITY_JSON_PATH/BridgeRegistry.json | jq -r ".networks[\"$ETHEREUM_NETWORK_ID\"].address")
export BRIDGE_TOKEN_ADDRESS=$(cat $SOLIDITY_JSON_PATH/BridgeToken.json | jq -r ".networks[\"$ETHEREUM_NETWORK_ID\"].address")
export BRIDGE_BANK_ADDRESS=$(cat $SOLIDITY_JSON_PATH/BridgeBank.json | jq -r ".networks[\"$ETHEREUM_NETWORK_ID\"].address")

cp $BASEDIR/smart-contracts/build/contracts/SojahubTestToken.json $SOLIDITY_JSON_PATH

echo ========== Sample commands ==========

echo; echo == efury balance
echo yarn -s --cwd $BASEDIR/smart-contracts integrationtest:getTokenBalance \
  --symbol \$BRIDGE_TOKEN_ADDRESS \
  --ethereum_private_key_env_var "ETHEREUM_PRIVATE_KEY" \
  --json_path \$BASEDIR/smart-contracts/deployments/$DEPLOYMENT_NAME \
  --gas estimate \
  --ethereum_network \$ETHEREUM_NETWORK \
  --ethereum_address \$ETHEREUM_ADDRESS \

echo; echo == eth balance
echo yarn -s --cwd $BASEDIR/smart-contracts integrationtest:getTokenBalance \
  --symbol eth \
  --ethereum_private_key_env_var "ETHEREUM_PRIVATE_KEY" \
  --json_path $BASEDIR/smart-contracts/deployments/$DEPLOYMENT_NAME \
  --gas estimate \
  --ethereum_network \$ETHEREUM_NETWORK \
  --ethereum_address \$ETHEREUM_ADDRESS \

echo; echo == mint efury
echo yarn -s --cwd /home/james/workspace/sojanode/smart-contracts integrationtest:mintTestnetTokens  \
  --symbol $BRIDGE_TOKEN_ADDRESS \
  --ethereum_private_key_env_var "OPERATOR_PRIVATE_KEY" \
  --json_path $BASEDIR/smart-contracts/deployments/$DEPLOYMENT_NAME \
  --gas estimate \
  --ethereum_network $ETHEREUM_NETWORK \
  --ethereum_address $ETHEREUM_ADDRESS \
  --operator_address $OPERATOR_ADDRESS \
  --amount 100000000000000000000000000

echo; echo == lock eth
echo yarn -s --cwd $BASEDIR/smart-contracts integrationtest:sendLockTx --sojahub_address $FURY_SOURCE \
  --symbol eth \
  --ethereum_private_key_env_var "ETHEREUM_PRIVATE_KEY" \
  --json_path $BASEDIR/smart-contracts/deployments/$DEPLOYMENT_NAME \
  --gas estimate \
  --ethereum_network $ETHEREUM_NETWORK \
  --bridgebank_address $BRIDGE_BANK_ADDRESS \
  --ethereum_address $ETHEREUM_ADDRESS \
  --amount 1700000000000000000

echo; echo == burn efury
echo yarn -s --cwd $BASEDIR/smart-contracts integrationtest:sendBurnTx \
  --symbol $BRIDGE_TOKEN_ADDRESS \
  --ethereum_private_key_env_var "ETHEREUM_PRIVATE_KEY" \
  --json_path $BASEDIR/smart-contracts/deployments/$DEPLOYMENT_NAME \
  --gas estimate \
  --ethereum_network $ETHEREUM_NETWORK \
  --bridgebank_address $BRIDGE_BANK_ADDRESS \
  --ethereum_address $ETHEREUM_ADDRESS \
  --sojahub_address $FURY_SOURCE \
  --amount 17

echo; echo == burn efury from operator account
echo yarn -s --cwd /home/james/workspace/sojanode/smart-contracts integrationtest:sendBurnTx \
  --symbol $BRIDGE_TOKEN_ADDRESS \
  --ethereum_private_key_env_var "OPERATOR_PRIVATE_KEY" \
  --json_path $BASEDIR/smart-contracts/deployments/$DEPLOYMENT_NAME \
  --gas estimate \
  --ethereum_network $ETHEREUM_NETWORK \
  --bridgebank_address $BRIDGE_BANK_ADDRESS \
  --ethereum_address $OPERATOR_ADDRESS \
  --sojahub_address $FURY_SOURCE \
  --amount 100000000000000000000000000

echo; echo == whitelisted tokens
echo yarn -s --cwd $BASEDIR/smart-contracts \
  integrationtest:whitelistedTokens \
  --bridgebank_address $BRIDGE_BANK_ADDRESS \
  --json_path $BASEDIR/smart-contracts/deployments/$DEPLOYMENT_NAME \
  --ethereum_network $ETHEREUM_NETWORK \

sojanodecmd=sojanoded

echo; echo == sojahub balance
echo $sojanodecmd q auth account --node $SOJANODE $FURY_SOURCE

echo; echo == sojahub transaction
echo $sojanodecmd q tx --node $SOJANODE --chain-id $DEPLOYMENT_NAME 193EFB4A5D20BEC58ADE8BACEB38264870ADD8BAFEA9D6DAABE554B0ACBC0C93

echo; echo == all account balances
echo "$sojanodecmd keys list --keyring-backend test --output json | jq -r '.[].address' | parallel $sojanodecmd q auth account --node $SOJANODE -o json {} | grep coins"

echo; echo == burn ceth
echo $sojanodecmd tx ethbridge burn \
  $FURY_SOURCE $ETHEREUM_ADDRESS 100 ceth 58560000000000000 \
  --node $SOJANODE \
  --keyring-backend test \
  --fees 100000fury \
  --ethereum-chain-id=$ETHEREUM_NETWORK_ID \
  --chain-id=$DEPLOYMENT_NAME  \
  --yes \
  --from $FURY_SOURCE \

echo; echo == send ceth
echo $sojanodecmd tx send $FURY_SOURCE sojasomedestination 100fury \
  --node $SOJANODE \
  --keyring-backend test \
  --fees 100000fury \
  --chain-id=$DEPLOYMENT_NAME  \
  --yes \

echo; echo == Simple test run against $DEPLOYMENT_NAME:
echo python3 -m pytest --color=yes -x -olog_cli=true -olog_level=DEBUG -v -olog_file=vagrant/data/pytest.log -v src/py/test_eth_transfers.py

echo; echo == Load test run against $DEPLOYMENT_NAME - change NTRANSFERS to a large number:
echo TOKENS=ceth,fury NTRANSFERS=2 python3 -m pytest -olog_level=DEBUG -olog_file=vagrant/data/pytest.log -v src/py/test_bulk_transfers_to_ethereum.py::test_bulk_transfers_from_sojahub

echo; echo
