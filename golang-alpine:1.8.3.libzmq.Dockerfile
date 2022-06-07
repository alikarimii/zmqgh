FROM golang:1.18.3-alpine
RUN apk --no-cache add gcc musl-dev libzmq zeromq-dev
