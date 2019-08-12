// Package kafka_producer kafka 生产者的包装
package kafka

import (
	"logcat/logcatch/constant"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/Shopify/sarama"
)

// Config 配置
type Config struct {
	Topic      string `xml:"topic"`
	Broker     string `xml:"broker"`
	Frequency  int    `xml:"frequency"`
	MaxMessage int    `xml:"max_message"`
}

type Producer struct {
	producer  sarama.AsyncProducer
	topic     string
	MsgQ      chan *sarama.ProducerMessage
	wg        sync.WaitGroup
	closeChan chan struct{}
}

// NewProducer 构造KafkaProducer
func NewProducer() (*Producer, error) {

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.NoResponse                                                          // Only wait for the leader to ack
	config.Producer.Compression = sarama.CompressionSnappy                                                    // Compress messages
	config.Producer.Flush.Frequency = time.Duration(constant.CON_KAKFA_PRODUCER_FREQUENCY) * time.Millisecond // Flush batches every 500ms
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	p, err := sarama.NewAsyncProducer(strings.Split(constant.CON_KAKFA_PRODUCER_BROKER, ","), config)
	if err != nil {
		return nil, err
	}
	ret := &Producer{
		producer:  p,
		topic:     constant.CON_KAKFA_PRODUCER_TOPICS,
		MsgQ:      make(chan *sarama.ProducerMessage, constant.CON_KAKFA_PRODUCER_MAXMESSAGE),
		closeChan: make(chan struct{}),
	}

	return ret, nil
}

// Run 运行
func (p *Producer) Run() {

	p.wg.Add(1)
	go func() {
		defer p.wg.Done()

	LOOP:
		for {
			select {
			case m := <-p.MsgQ:
				p.producer.Input() <- m
			case err := <-p.producer.Errors():
				if nil != err && nil != err.Msg {
					log.Fatalf("[producer] err=[%s] topic=[%s] key=[%s] val=[%s]", err.Error(), err.Msg.Topic, err.Msg.Key, err.Msg.Value)
				}
			case <-p.closeChan:
				break LOOP
			}

		}
	}()

	for hasTask := true; hasTask; {
		select {
		case m := <-p.MsgQ:
			p.producer.Input() <- m
		default:
			hasTask = false
		}
	}

}

// Close 关闭
func (p *Producer) Close() error {
	close(p.closeChan)
	log.Fatalf("[producer] is quiting")
	p.wg.Wait()
	log.Fatalf("[producer] quit over")

	return p.producer.Close()
}

// Log 发送log
func (p *Producer) Log(key string, val string, partition int32) {
	msg := &sarama.ProducerMessage{
		Topic:     p.topic,
		Key:       sarama.StringEncoder(key),
		Value:     sarama.StringEncoder(val),
		Partition: partition,
	}

	select {
	case p.MsgQ <- msg:
		return
	default:
		log.Fatalf("[producer] err=[msgQ is full] key=[%s] val=[%s]", msg.Key, msg.Value)
	}
}
