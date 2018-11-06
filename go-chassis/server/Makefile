shell=/bin/bash

all:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server
	./build.sh
	distribute-image.sh xiaoliang/server-chassis:latest
	kubectl delete -f ./server-chassis.yaml
	kubectl apply -f ./server-chassis.yaml
