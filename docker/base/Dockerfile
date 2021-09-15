FROM alpine:3.14

ENV SCALA_VER=2.13 KAFKA_VER=2.8.0 PATH=/kafka/bin:$PATH

RUN apk update && apk add ca-certificates bash openjdk11

RUN wget https://dlcdn.apache.org/kafka/$KAFKA_VER/kafka_$SCALA_VER-$KAFKA_VER.tgz  && tar -xzf kafka_$SCALA_VER-$KAFKA_VER.tgz && mv kafka_$SCALA_VER-$KAFKA_VER kafka  && rm kafka_$SCALA_VER-$KAFKA_VER.tgz