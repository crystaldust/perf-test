FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY client /root/client
COPY conf /root/conf
WORKDIR /root/
ADD data ./data
ENTRYPOINT ["./client"]
