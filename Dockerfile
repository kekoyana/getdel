FROM golang:1.15-alpine3.13

RUN apk add git
RUN go env -w GO111MODULE=on

RUN mkdir /work
WORKDIR /work

ADD . /work
