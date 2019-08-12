package service

import (
	"logcat/consumer/httpsvr"
	"logcat/consumer/service/elastic"
	"logcat/consumer/service/kafka"
)

type Service struct {
	Elastic *elastic.ElasticManager
	Kafka   *kafka.KafkaManager
	Httpsvr *httpsvr.HttpServer
}

func NewService() *Service {

	//ES 或者 HBASE
	es := elastic.NewEsManager()

	//ka := kafka.NewKafkaManager()

	httpsvr := httpsvr.NewHttpServer(nil)

	return &Service{
		Elastic: es,
		Kafka:   nil,
		Httpsvr: httpsvr,
	}
}

func (this *Service) Start() {

	//runtime.GOMAXPROCS(runtime.NumCPU()) // 利用CPU多核来处理HTTP请求

	this.Httpsvr.Start()

}
