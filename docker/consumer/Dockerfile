FROM kafka-base

ENV TOPIC=email-notification

RUN adduser -D consumer

USER consumer

WORKDIR /home/consumer

CMD ["sh", "-c", "kafka-console-consumer.sh --topic $TOPIC --bootstrap-server broker:9092"]