ARG BASE_IMAGE="1.19-alpine"

FROM golang:$BASE_IMAGE

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
COPY --chown=test:testgroup go.mod ./


RUN go mod download

WORKDIR /app/test

ENTRYPOINT [ "task", "test-inside-docker" ]