#!/bin/bash

function test() {
	addr=$1
	path=$2
	wrk -t20 -c40 -d10s http://$addr/$path
}

sleep 30
test localhost:8000 json
sleep 30
test localhost:8000 image
echo

sleep 30
test localhost:8001 json
sleep 30
test localhost:8001 image
echo

sleep 30
test localhost:8002 json
sleep 30
test localhost:8002 image
echo
