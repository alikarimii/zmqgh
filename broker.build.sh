#!/bin/bash
if [ ! $@ ]; then
name='broker-img'
else
name=$@
fi
docker build -t $name \
-f broker.Dockerfile .