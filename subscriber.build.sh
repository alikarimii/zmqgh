#!/bin/bash
if [ ! $@ ]; then
name='subscriber-img'
else
name=$@
fi
docker build -t $name \
-f subscriber.Dockerfile .