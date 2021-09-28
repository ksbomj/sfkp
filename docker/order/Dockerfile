FROM alpine:latest as base

RUN apk add go ca-certificates git

COPY ./services/order/ /order

WORKDIR /order

RUN ls -la && go mod vendor && go build .

FROM alpine:latest

RUN apk add ca-certificates curl

COPY --from=base /order/order /bin

RUN adduser -D order-service

USER order-service

CMD ["/bin/order", "server"]


