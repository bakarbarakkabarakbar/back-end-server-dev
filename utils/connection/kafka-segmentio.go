package connection

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

type BrokerSegmentIo struct {
	Host      string
	Port      string
	Topic     string
	Partition int
}

func NewBrokerSegmentIo(host string, port string, topic string, partition int) BrokerSegmentIo {
	return BrokerSegmentIo{
		Host:      host,
		Port:      port,
		Topic:     topic,
		Partition: partition,
	}
}

func (b BrokerSegmentIo) WriteMessage(key string, msg string) error {
	var err error
	var url = fmt.Sprintf("%s:%s", b.Host, b.Port)
	var leader *kafka.Conn
	leader, err = kafka.DialLeader(context.Background(), "tcp", url, b.Topic, b.Partition)
	if err != nil {
		return err
	}
	var messages int
	messages, err = leader.WriteMessages(kafka.Message{Key: []byte(key), Value: []byte(msg)})
	if err != nil {
		return err
	}
	fmt.Println(messages)
	err = leader.Close()
	if err != nil {
		return err
	}
	return nil
}

func (b BrokerSegmentIo) GetMessages() error {
	var err error
	var url = fmt.Sprintf("%s:%s", b.Host, b.Port)
	var leader *kafka.Conn
	leader, err = kafka.DialLeader(context.Background(), "tcp", url, b.Topic, b.Partition)
	if err != nil {
		return err
	}

	leader.SetReadDeadline(time.Now().Add(10 * time.Second))
	var batchRead = leader.ReadBatch(0, 1e6)

	var message kafka.Message
	for {
		message, err = batchRead.ReadMessage()
		if err != nil {
			break
		}
		fmt.Printf("Offset: %d, Key: %s, Message: %s\n", message.Offset, message.Key, message.Value)

	}

	err = batchRead.Close()
	if err != nil {
		return err
	}
	err = leader.Close()
	if err != nil {
		return err
	}
	return nil
}
