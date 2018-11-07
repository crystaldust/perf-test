#!/usr/bin/env bash
cp ../../testdata/* ./
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o client .
sudo docker build -t perf-test/go-restful-client ./
rm sample.*
