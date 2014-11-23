#!/bin/bash

mkdir /tmp/roj-consul
consul agent --server --bootstrap --data-dir=/tmp/roj-consul
rm -rf /tmp/roj-consul 