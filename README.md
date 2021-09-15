## Scripts for Kafka Presentation

Tested on Linux only, have no idea how it's going to work on Windows.

My setup is build on top of docker because I didn't want to pollute my machine with Java and Kafka installations.

### Binary dependencies

* docker
* docker-compose

## Test

```posix
git clone https://github.com/ksbomj/sfkp
cd sfkp
./setup.sh
```

These commands will start containers for **zookeeper**, **kafka broker**, **producer** and **consumers**. For producer I have one container that publish into 5 different topics, and 5 separate containers for consumers. I did so because I didn't want to overload compose file with producer for each topic, however it's useful to see different colors output for consumers in setup.sh script.

## Schema

Didn't spend a lot of time on schema because of few reasons:

1. it's not enough context information about shop domain to specify event body


2. I didn't get why I need to use an "Avro" schema for Kafka (I know it's optional) and why they invented it. I'm looking into something like gob, messagepack or at least protobuf to use for future.
	
Schema files are located [here](https://github.com/ksbomj/sfkp/tree/main/docker/producer/events-schema).
