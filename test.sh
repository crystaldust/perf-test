#!/bin/bash

function test() {
	addr=$1
	path=$2
	wrk -t20 -c40 -d10s http://$addr/$path
}

function batch() {
	for i in {1..3}
	do
		sleep 30
		test $1 $2
	done
}


batch localhost:8000 json
batch localhost:8000 image

batch localhost:8001 json
batch localhost:8001 image

batch localhost:8002 json
batch localhost:8002 image
