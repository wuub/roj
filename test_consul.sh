#!/bin/bash

mkdir /tmp/roj-consul
consul agent --server --bootstrap --data-dir=/tmp/roj-consul -advertise=127.0.0.1
rm -rf /tmp/roj-consul 