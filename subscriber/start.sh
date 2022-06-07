#!/bin/bash
if [ ! $@ ]; then
name='subscriber-ins'
else
name=$@
fi
docker rmi `docker ps -aq -f name=$name`
set -a
source .env

docker run --rm \
--network zmqgh-network \
-e SUBSCRIBER_REP_ENDPOINT=$SUBSCRIBER_REP_ENDPOINT \
--name $name subscriber-img

# background
# docker run -d --rm \
# --network zmqgh-network \
# -e SUBSCRIBER_REP_ENDPOINT=$SUBSCRIBER_REP_ENDPOINT \
# --name $name subscriber-img