shell=/bin/bash

all:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o client
	./build.sh
	distribute-image.sh xiaoliang/client-chassis:latest
	kubectl delete -f ./client-chassis.yaml
	kubectl apply -f ./client-chassis.yaml
