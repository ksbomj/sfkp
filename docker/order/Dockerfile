FROM alpine:latest as base

RUN apk add go ca-certificates git librdkafka librdkafka-dev musl musl-dev gcc g++

COPY ./services/order/ /order

WORKDIR /order

RUN ls -la && go mod vendor && CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -v -tags musl .

FROM alpine:latest

RUN apk add ca-certificates curl

COPY --from=base /order/order /bin

RUN adduser -D order-service

USER order-service

CMD ["/bin/order", "server"]


