FROM kafka-base

RUN adduser -D producer

USER producer

WORKDIR /home/producer

COPY producer.sh .
COPY --chown=producer  ./events-schema  events-schema

CMD ["./producer.sh"]