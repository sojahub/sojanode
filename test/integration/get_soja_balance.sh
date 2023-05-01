#!/bin/bash 

addr=$1
shift

sojanoded q auth account ${addr:=${VALIDATOR1_ADDR}} -o json | jq
