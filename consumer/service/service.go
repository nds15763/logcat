package service

import (
	"careyads/esserver/httpsvr"
	"careyads/esserver/service/elastic"
	"careyads/esserver/service/kafka"
)

type EsService struct {
	Elastic *elastic.ElasticManager
	Kafka   *kafka.KafkaManager
	Httpsvr *httpsvr.HttpServer
}

func NewService() *EsService {

	//es := elastic.NewEsManager()

	ka := kafka.NewKafkaManager()

	httpsvr := httpsvr.NewHttpServer(nil)

	return &EsService{
		Elastic: nil,
		Kafka:   ka,
		Httpsvr: httpsvr,
	}
}

func (this *EsService) Start() {

	//runtime.GOMAXPROCS(runtime.NumCPU()) // 利用CPU多核来处理HTTP请求

	this.Httpsvr.Start()

}
