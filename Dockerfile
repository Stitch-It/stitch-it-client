FROM alpine:latest

RUN apk add --no-cache git go

ARG BEARER_TOKEN=""

ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH
ENV BEARER_TOKEN=$BEARER_TOKEN

WORKDIR /go/src/twitter-bot
COPY . .
ENV GO111MODULE=on
RUN go build .
CMD [ "./stitch-it" ]
