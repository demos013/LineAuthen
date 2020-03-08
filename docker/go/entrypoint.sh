#!/bin/bash

cd /go/src/healthcheck

echo "Install dependency..."
go get ./...

echo "Start service..."
exec gin -p 8000 -a 8080 "$@"