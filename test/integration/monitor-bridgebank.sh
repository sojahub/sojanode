#!/bin/bash 

# usually run like this:
# nohup bash /sojanode/test/integration/monitor-bridgebank.sh > /tmp/bridgebank.txt &

while true
do
  date
  docker exec integration_sojanode1_1 bash -c "cd /smart-contracts && yarn peggy:getTokenBalance $BRIDGE_BANK_ADDRESS eth" >> /tmp/bridgebank.txt
  sleep 5
done