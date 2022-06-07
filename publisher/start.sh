#!/bin/bash
if [ ! $@ ]; then
name='publisher-ins'
else
name=$@
fi
docker rmi `docker ps -aq -f name=$name`
set -a
source .env

docker run --rm \
--network zmqgh-network \
-e SOURCE_REQ_ENDPOINT=$SOURCE_REQ_ENDPOINT \
-e REQUEST_TIMEOUT=$REQUEST_TIMEOUT \
-e REQUEST_RETRIES=$REQUEST_RETRIES \
--name $name publisher-img