#!/usr/bin/env bash
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .
sudo docker build -t perf-test/go-chassis-server ./
