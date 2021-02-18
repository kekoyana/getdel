FROM golang:1.15-alpine3.13

RUN mkdir /work
WORKDIR /work

ADD . /work

