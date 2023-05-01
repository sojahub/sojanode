#!/usr/bin/env bash

echo "Creating pools ceth and cdash"
sojanoded tx clp create-pool --from soja --symbol ceth --nativeAmount 20000000000000000000 --externalAmount 20000000000000000000  --yes

sleep 5
sojanoded tx clp create-pool --from soja --symbol cdash --nativeAmount 20000000000000000000 --externalAmount 20000000000000000000  --yes


sleep 8
echo "Swap Native for Pegged - Sent fury Get ceth"
sojanoded tx clp swap --from soja --sentSymbol fury --receivedSymbol ceth --sentAmount 2000000000000000000 --minReceivingAmount 0 --yes
sleep 8
echo "Swap Pegged for Native - Sent ceth Get fury"
sojanoded tx clp swap --from soja --sentSymbol ceth --receivedSymbol fury --sentAmount 2000000000000000000 --minReceivingAmount 0 --yes
sleep 8
echo "Swap Pegged for Pegged - Sent ceth Get cdash"
sojanoded tx clp swap --from soja --sentSymbol ceth --receivedSymbol cdash --sentAmount 2000000000000000000 --minReceivingAmount 0 --yes

sojanoded q clp pools

