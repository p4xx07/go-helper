package consumer

import (
	"fmt"
	"github.com/Shopify/sarama"
)

type IConsumer interface {
	Consume(topic string, offset int64, onConsume func(msg *sarama.ConsumerMessage) error, maxThreads int) error
}

type consumer struct {
	saramaConsumer sarama.Consumer
}

func NewConsumer(brokers []string, offset int64, configParams ...sarama.Config) (IConsumer, error) {
	config := prepareConfig(configParams, offset)
	saramaConsumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &consumer{saramaConsumer: saramaConsumer}, nil
}

func prepareConfig(configParams []sarama.Config, offset int64) *sarama.Config {
	if len(configParams) > 0 {
		return &configParams[0]
	}
	return getDefaultConfig(offset)
}

func getDefaultConfig(offset int64) *sarama.Config {
	var config *sarama.Config
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = offset
	return config
}

func (c *consumer) Consume(topic string, offset int64, onConsume func(msg *sarama.ConsumerMessage) error, maxThreads int) error {
	partitions, err := c.saramaConsumer.Partitions(topic)
	if err != nil {
		return err
	}

	maxThreadChannel := make(chan bool, maxThreads)

	for _, partition := range partitions {
		partitionConsumer, err := c.saramaConsumer.ConsumePartition(topic, partition, offset)

		if err != nil {
			return err
		}

		go func() {
			defer partitionConsumer.Close()
			for {
				select {
				case msg := <-partitionConsumer.Messages():
					maxThreadChannel <- true
					go func() {
						_ = onConsume(msg)
						<-maxThreadChannel
					}()

				case err = <-partitionConsumer.Errors():
					fmt.Println(err.Error())
				}
			}
		}()
	}

	return nil
}
