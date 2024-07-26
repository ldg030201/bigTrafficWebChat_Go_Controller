package kafka

import (
	"chat_controller_server/config"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Kafka struct {
	cfg *config.Config

	consumer *kafka.Consumer
}

func NewKafka(cfg *config.Config) (*Kafka, error) {
	k := &Kafka{cfg: cfg}

	var err error

	if k.consumer, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.Kafka.URL,
		"group.id":          cfg.Kafka.GroupID,
		"auto.offset.reset": "latest",
	}); err != nil {
		return nil, err
	} else {
		return k, nil
	}
}

func (k *Kafka) RegisterSubTopic(topic string) error {
	if err := k.consumer.Subscribe(topic, nil); err != nil {
		return err
	} else {
		return nil
	}
}

func (k *Kafka) Pool(time int) kafka.Event {
	return k.consumer.Poll(time)
}
