FROM alpine:latest
RUN apk --no-cache add ca-certificates
ADD server /root/
ADD conf /root/conf
ADD start.sh /root/start.sh
RUN chmod +x /root/start.sh
WORKDIR /root/
CMD ["./start.sh"]
