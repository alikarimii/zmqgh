FROM alpine:3
RUN apk --no-cache add gcc musl-dev libzmq zeromq-dev
