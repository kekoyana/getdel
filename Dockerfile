FROM golang:1.15-alpine3.13

RUN apk add --no-cache gcc musl-dev git sqlite
RUN go env -w GO111MODULE=on

RUN mkdir /work
WORKDIR /work

ADD . /work
