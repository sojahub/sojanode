#!/usr/bin/env bash

killall sojanoded

sojanoded start --trace --home ~/.sojanoded/ --p2p.laddr 0.0.0.0:27655  --grpc.address 0.0.0.0:9096 --grpc-web.address 0.0.0.0:9093 --address tcp://0.0.0.0:27659 --rpc.laddr tcp://127.0.0.1:26657
