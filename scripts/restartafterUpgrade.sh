#!/usr/bin/env bash

cp $GOPATH/src/new/sojanoded $GOPATH/bin/
cosmovisor start >> sojanode.log 2>&1  &