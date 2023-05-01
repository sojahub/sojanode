#!/bin/bash
pid_sojanode=$(ps aux | grep "sojanoded start" | grep -v grep | awk '{print $2}')
# pid_rest=$(ps aux | grep "sojanoded rest-server" | grep -v grep | awk '{print $2}')

if [[ ! -z "$pid_sojanode" ]]; then 
  kill -9 $pid_sojanode
fi

# if [[ ! -z "$pid_rest" ]]; then 
#   kill -9 $pid_rest
# fi