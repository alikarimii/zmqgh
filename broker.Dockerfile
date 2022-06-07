FROM golang-alpine-libzmq:1.8.3 as build-env

WORKDIR $GOPATH/zmqgh/broker
COPY go.mod ./go.mod
COPY go.sum ./go.sum
RUN go mod download

COPY broker ./broker
COPY pkg ./pkg
# COPY pb ./pb if have

RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o goapp ./broker/cmd

FROM alpine.libzmq:3
RUN mkdir /app
# Create user and set ownership and permissions as required
RUN adduser -D myuser && chown -R myuser /app
WORKDIR /app
USER myuser
COPY --from=build-env /go/zmqgh/broker/goapp .

EXPOSE 40001
ENTRYPOINT ["./goapp"]