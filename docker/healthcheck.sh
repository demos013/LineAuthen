#!/bin/sh


docker rm -f healthcheck
# docker-compose build coreservice
# docker-compose run -d -p 8008:8008 --name coreservice coreservice main.go

docker-compose  up -d --build healthcheck