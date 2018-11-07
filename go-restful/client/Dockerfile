FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY client ./client
COPY sample.png ./sample.png
COPY sample.json ./sample.json
ENTRYPOINT ["./client"]
