FROM golang:1.19-alpine

LABEL maintainer="iscoct@correo.ugr.es"

WORKDIR /app

RUN addgroup -S testgroup && adduser -S -u 1001 test -G testgroup
RUN chown test /app
RUN apk add build-base

USER test

RUN mkdir test
RUN mkdir ejecucion_tests

WORKDIR /app/ejecucion_tests

RUN go install github.com/go-task/task/v3/cmd/task@latest
COPY go.mod ./

USER root

RUN chown test go.mod

USER test

RUN go mod download

WORKDIR /app/test

ENTRYPOINT [ "task", "test-inside-docker" ]