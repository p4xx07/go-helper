package producer

import (
	"bytes"
	"encoding/json"
	"github.com/Shopify/sarama"
)

type IService interface {
	ProduceFromString(topic string, message string) error
	ProduceFromAny(topic string, payload any) error
	Close() error
}

type service struct {
	producer sarama.SyncProducer
}

func NewService(brokers []string, configParams ...sarama.Config) (IService, error) {
	config := prepareConfig(configParams)
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &service{producer: producer}, nil
}

func prepareConfig(configParams []sarama.Config) *sarama.Config {
	if len(configParams) > 0 {
		return &configParams[0]
	}
	return getDefaultConfig()
}

func getDefaultConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message
	config.Producer.Retry.Max = 10                   // Retry up to 10 times to produce the message
	config.Producer.Return.Successes = true
	return config
}

func (p *service) ProduceFromString(topic string, message string) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	_, _, err := p.producer.SendMessage(msg)
	if err != nil {
		return err
	}

	return nil
}

func (p *service) ProduceFromAny(topic string, payload any) error {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)

	if err := enc.Encode(payload); err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(buf.Bytes()),
	}

	_, _, err := p.producer.SendMessage(msg)
	if err != nil {
		return err
	}
	return nil
}

func (p *service) Close() error {
	return p.producer.Close()
}
