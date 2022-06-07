# Hexagonal Architecture Practicing

## Run in docker environment
1. first create base docker image
```
docker build -t golang-alpine-libzmq:1.8.3 -f golang-alpine:1.8.3.libzmq.Dockerfile .
docker build -t alpine.libzmq:3 -f alpine.libzmq.Dockerfile .
```
2. set your config in `.env` file.
```
.
|-- broker
|   |-- ...
|   |-- .env
|   `-- start.sh
|
|-- publisher
|   |-- ...
|   |-- .env
|   `-- start.sh
|
|-- subscriber
|   |-- ...
|   |-- .env
|   `-- start.sh
|
|
|-- README.md
|-- ...
|-- go.mod
|-- go.sum
```
3. set +X permission to `.sh` file
```
chmod 710 broker.build.sh
chmod 710 publisher.build.sh
chmod 710 subscriber.build.sh
chmod 710 broker/start.sh
chmod 710 publisher/start.sh
chmod 710 subscriber/start.sh
```
4. build images
```
./broker.build.sh
./publisher.build.sh
./subscriber.build.sh
```
5. create `zmqgh-network` network
```
docker network create zmqgh-network
```

6. run containers.you must change dir to each folder and exec start.`cd broker`
```
./start.sh
```
- to run each container in background add `-d` to docker run command in `start.sh`
