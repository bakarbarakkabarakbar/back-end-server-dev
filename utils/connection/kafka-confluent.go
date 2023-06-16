package connection

//
//import (
//	"fmt"
//	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
//)
//
//type BrokerConfluent struct {
//	Host string
//	Port string
//}
//
//type ConsumerConfluent struct {
//	broker BrokerConfluent
//}
//
//func NewBrokerConfluent(host string, port string) BrokerConfluent {
//	return BrokerConfluent{
//		Host: host,
//		Port: port,
//	}
//}
//
//func NewConsumerConfluent(broker BrokerConfluent) ConsumerConfluent {
//	return ConsumerConfluent{broker: broker}
//}
//
//func (b BrokerConfluent) Init() (*kafka.Producer, error) {
//	var url = fmt.Sprintf("%s:%s", b.Host, b.Port)
//	var producer, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": url})
//	if err != nil {
//		return nil, err
//	}
//	return producer, nil
//}
//
//func (c ConsumerConfluent) Init() (*kafka.Consumer, error) {
//	var url = fmt.Sprintf("%s:%s", c.broker.Host, c.broker.Port)
//	var consumer, err = kafka.NewConsumer(&kafka.ConfigMap{
//		"bootstrap.servers":  url,
//		"group.id":           "my-consumer-group",
//		"auto.offset.reset":  "earliest",
//		"enable.auto.commit": "false",
//	})
//	if err != nil {
//		return nil, err
//	}
//	return consumer, nil
//}
