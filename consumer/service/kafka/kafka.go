package kafka

import (
	"logcat/consumer/constant"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	"github.com/bsm/sarama-cluster"
)

type KafkaManager struct {
	Consumer *cluster.Consumer
}

func NewKafkaManager() *KafkaManager {

	config := cluster.NewConfig()
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	config.Consumer.Offsets.Initial = sarama.OffsetNewest //初始从最新的offset开始

	consumer, err := cluster.NewConsumer(strings.Split(constant.CON_KAKFA_ADRRESS1, ","), constant.CON_KAKFA_GROUPID, strings.Split(constant.CON_KAKFA_TOPICS, ","), config)
	if err != nil {
		//logrus.Errorf("Failed open consumer: %s", err.Error())
		log.Fatalf("Error getting response: %s", err)
		return nil
	}

	return &KafkaManager{
		Consumer: consumer,
	}

}

func (this *KafkaManager) Start() {

	// consume messages, watch signals
	var successes int
	for {
		select {
		case msg, ok := <-this.Consumer.Messages():
			if ok {
				fmt.Fprintf(os.Stdout, "%s:%s/%s/%s\t%s\t%s\n", constant.CON_KAKFA_GROUPID, constant.CON_KAKFA_TOPICS, "msg.Partition", "msg.Offset", "msg.Key", "msg.Value")
				this.Consumer.MarkOffset(msg, "") // mark message as processed
				successes++
			}
		}
	}

}
