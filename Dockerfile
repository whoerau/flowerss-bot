FROM golang:alpine as builder

COPY . /flowerss
RUN apk add git make gcc libc-dev && \
    cd /flowerss && make build


FROM alpine
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /flowerss/flowerss-bot /bin/
VOLUME /root/.flowerss
WORKDIR /root/.flowerss
ENTRYPOINT ["/bin/flowerss-bot"]

