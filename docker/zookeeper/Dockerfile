FROM kafka-base

RUN adduser -D zookeeper && chown -R zookeeper:zookeeper /kafka

USER zookeeper

WORKDIR /home/zookeeper

COPY --chown=zookeeper config config

CMD ["zookeeper-server-start.sh", "config/zookeeper.properties"]