FROM kafka-base

RUN adduser -D broker && chown -R broker:broker /kafka

USER broker

WORKDIR /home/broker

COPY --chown=broker config config

CMD ["kafka-server-start.sh", "config/server.properties"]