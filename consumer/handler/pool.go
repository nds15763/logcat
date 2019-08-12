package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"logcat/consumer/constant"
	"logcat/logcatch/model"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

//map每一个KEY就相当于一个group，每个group就只消费一个topic，每一个chan有partition个容量，就相当于pool

//而对于最后的map，消费端只是希望能获取到一个队列

//并不涉及顺序问题，如果涉及排序问题需要指定partition，让每个consumer顺序执行

type KafkaMap struct {
	topics []string
	kaMap  map[string]chan *model.InLog
}

func NewKafkaMap(allTopic []string) *KafkaMap {

	tmp := make(map[string]chan *model.InLog, 0)

	for _, v := range allTopic {
		tmp[v] = make(chan *model.InLog, 3)
	}

	return &KafkaMap{
		topics: allTopic,
		kaMap:  tmp,
	}
}

func (this *KafkaMap) Run() {
	//遍历map,启动consumer
	i := 1
	for k, _ := range this.kaMap {
		//新建consumer,key为topic字段
		c, err := NewConsumer(i, k)
		if err != nil {
			log.Fatalf("Error NewConsumer: %s", err)
		}
		defer c.Close()

		//持久化consumer
		go func(c *cluster.Consumer) {
			errors := c.Errors()
			noti := c.Notifications()
			for {
				select {
				case err := <-errors:
					log.Fatalf("Error Consumer Notifications: %s", err)
				case <-noti:
				}
			}
		}(c)

		go func(c *cluster.Consumer) {
			for msg := range c.Messages() {
				inlog := &model.InLog{}

				err = json.Unmarshal(msg.Value, inlog)
				if err != nil {
					log.Fatalf("Consumer Json.marshal: %s", err)
					return
				}

				this.kaMap[msg.Topic] <- inlog
				c.MarkOffset(msg, "")
			}
		}(c)

		i++

	}
}

func NewConsumer(group int, topic string) (*cluster.Consumer, error) {
	config := cluster.NewConfig()
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	config.Consumer.Offsets.Initial = sarama.OffsetNewest //初始从最新的offset开始
	//config

	consumer, err := cluster.NewConsumer(strings.Split(constant.CON_KAKFA_ADRRESS1, ","), fmt.Sprintf("%d", group), strings.Split(constant.CON_KAKFA_TOPICS, ","), config)
	if err != nil {
		return nil, err
	}

	return consumer, nil
}

func doConsume(con *cluster.Consumer) {

}
