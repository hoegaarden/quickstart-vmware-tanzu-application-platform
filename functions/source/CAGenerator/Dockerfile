FROM golang:1.20-alpine

RUN apk --no-cache add make git zip

COPY . /build

WORKDIR /build

RUN GOPROXY=direct go mod download

RUN GOPROXY=direct make -f Makefile.package package

CMD mkdir -p /output/ && mv /build/CAGenerator.zip /output/
