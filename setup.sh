#!/bin/sh

RED='\033[0;31m'
GREEN='\033[0;32m'

DEPENDENCIES="docker docker-compose"

for dep in $DEPENDENCIES
do
  if [[ ! -x "$(command -v $dep)" ]] ; then
	  echo "${RED}Error: $dep is not installed"
	  exit 1
  fi
done

# Building a base image
docker build -t kafka-base -f ./docker/base/Dockerfile ./docker/base

# Build compose containers
docker-compose -f  kafka-docker-compose.yml -f ppe-docker-compose.yml build

# Run containers
docker-compose -f kafka-docker-compose.yml -f ppe-docker-compose.yml up -d --force-recreate

echo "\n${GREEN}Topics registration"

sleep 5

TOPICS="order-received order-confirmed order-picked-and-packed email-notification event-proccessing-error"

for topic in $TOPICS
do
  docker exec -t broker sh -c 'kafka-topics.sh --describe --topic '$topic' --bootstrap-server localhost:9092 ; exit $?'

  if [ $? -ne 0 ] ; then
    docker exec -t broker sh -c 'kafka-topics.sh --create --topic '$topic' --bootstrap-server localhost:9092 ; exit $?'

    if [ $? -ne 0 ] ; then
        echo "${RED}Error: coudn't create a topic \"$topic\""
        exit 1
    fi

    echo "\xE2\x9C\x94 Topic \"$topic\" was successfully created"
  else
    # Alter topic doesn't work fork for the version 2.8.0, see details https://issues.apache.org/jira/browse/KAFKA-8406
    #
    # docker exec -t broker sh -c 'kafka-topics.sh --alter --topic '$topic'  --config retention.ms=10800000 --bootstrap-server localhost:9092 ; exit $?'
    #
    # if [ $? -ne 0 ] ; then
    #   echo "${RED}Error: coudn't change retention.ms config option for  topic \"$topic\""
    #   exit 1
    # fi

    echo "\xE2\x9C\x94 Topic \"$topic\" was updated, retention time was increased to 3 days"
  fi
done

echo "\n${GREEN}Producing events"

docker logs producer

echo "\n${GREEN}Consuming events"

docker-compose -f  kafka-docker-compose.yml -f ppe-docker-compose.yml logs consumer1 consumer2 consumer3 consumer4 consumer5
