# Install the base requirements for the app.
# This stage is to support development.
FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/lkjfrf/client/main

EXPOSE 8080
ENTRYPOINT [ "/client/main.go" ]