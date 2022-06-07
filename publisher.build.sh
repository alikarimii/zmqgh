#!/bin/bash
if [ ! $@ ]; then
name='publisher-img'
else
name=$@
fi
docker build -t $name \
-f publisher.Dockerfile .