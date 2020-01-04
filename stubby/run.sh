#!/usr/bin/env bash
IP_ADDR="$( ifconfig | grep "inet " | grep -Fv 127.0.0.1 | awk '{print $2}' )"
java -jar stubby4j-6.0.2.jar -d run.yaml -l $IP_ADDR -s 8888 -w
