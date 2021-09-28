package bus

import (
	"encoding/json"
	"github.com/ksbomj/sfkp/services/order/events"
	"log"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type Bus struct {
	brokerAddress string
	socketTimeout int
	deliveryTimeout int
	logger *log.Logger
}

func NewBus(brokerAddress string, logger *log.Logger) *Bus {
	return &Bus{
		brokerAddress: brokerAddress,
		socketTimeout: 30000,
		logger: logger,
	}
}

func (b Bus) PublishEvent(event events.Event, topic string) error {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":   b.brokerAddress,
		"socket.timeout.ms":   b.socketTimeout,
		"delivery.timeout.ms": b.deliveryTimeout})

	if err != nil {
		return err
	}

	deliveryChan := make(chan kafka.Event)

	var value []byte
	if value, err = json.Marshal(event); err != nil {
		return err
	}
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          value,
	}, deliveryChan)

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		return m.TopicPartition.Error
	}

	b.logger.Printf(
		"Topic delivered.\n Name: %s\nPartition: %s\nPartitionOffset: %s\n",
		*m.TopicPartition.Topic,
		m.TopicPartition.Partition,
		m.TopicPartition.Offset,
	)

	close(deliveryChan)

	return nil
}
