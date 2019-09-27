FROM golang:alpine AS build-env

ADD . /go/src/github.com/github.com/dubuqingfeng/explorer_parser/src
# build
RUN cd /go/src/github.com/github.com/dubuqingfeng/explorer_parser/src && \
    go build -v -o /src/bin/consumer consumer/consumer.go && \
    go build -v -o /src/bin/producer producer/producer.go

FROM alpine
RUN apk --no-cache add openssl ca-certificates
WORKDIR /app
COPY --from=build-env /src/bin /app/
COPY --from=build-env /go/src/github.com/dubuqingfeng/explorer_parser/src/config /app/config
COPY --from=build-env /go/src/github.com/dubuqingfeng/explorer_parser/src/logs /app/logs
ENTRYPOINT ./notice-web