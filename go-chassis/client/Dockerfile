FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY client /root/client
COPY conf /root/conf
WORKDIR /root/
ADD sample.png ./sample.png
ADD sample.json ./sample.json
ENTRYPOINT ["./client"]
