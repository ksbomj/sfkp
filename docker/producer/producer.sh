#!/bin/sh

sleep 15
ls ./events-schema | while read -r file; do cat ./events-schema/$file | kafka-console-producer.sh --topic $(echo $file | cut -d  '.'  -f 1)  --bootstrap-server broker:9092; done
