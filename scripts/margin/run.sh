#!/usr/bin/env bash

set -x

killall sojanoded

cd ../..
make install
sojanoded start --trace
