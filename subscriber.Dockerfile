FROM golang-alpine-libzmq:1.8.3 as build-env

WORKDIR $GOPATH/zmqgh/subscriber
COPY go.mod ./go.mod
COPY go.sum ./go.sum
RUN apk --no-cache add gcc musl-dev libzmq zeromq-dev
RUN go mod download

COPY subscriber ./subscriber
COPY pkg ./pkg
# COPY pb ./pb if have

RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o goapp ./subscriber/cmd

FROM alpine.libzmq:3
RUN mkdir /app
# Create user and set ownership and permissions as required
RUN adduser -D myuser && chown -R myuser /app
WORKDIR /app
USER myuser
COPY --from=build-env /go/zmqgh/subscriber/goapp .

EXPOSE 40002
ENTRYPOINT ["./goapp"]