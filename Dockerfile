FROM alpine:latest as builder

RUN apk add --no-cache git go

ENV GOOS js
ENV GOARCH wasm
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH
ENV BEARER_TOKEN=$BEARER_TOKEN

WORKDIR /go/src/stitch-it
COPY ./go .
ENV GO111MODULE=on
RUN go build -o ./main.wasm ./main.go

# this is where the front end will get built
# copy the built WASM file from builder to
# this step, whatever image we decide to build it on
