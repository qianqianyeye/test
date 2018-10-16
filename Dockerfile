
FROM golang:latest

MAINTAINER Razil "996888231@qq.com"

WORKDIR $GOPATH/src/ServiceTest
ADD . $GOPATH/src/ServiceTest
RUN godep go build .

EXPOSE 8011

ENTRYPOINT ["./ServiceTest"]
