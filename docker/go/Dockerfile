FROM golang:1.14

RUN mkdir -p /go/src/healthcheck
WORKDIR /go/src/healthcheck

COPY entrypoint.sh /root/

RUN chmod +x /root/entrypoint.sh
ENV GOPATH /go/

RUN go get github.com/codegangsta/gin