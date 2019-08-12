package service

import (
	"logcat/logcatch/httpsvr"
	"logcat/logcatch/service/kafka"
	"runtime"
)

type Service struct {
	Kafka   *kafka.Producer
	Httpsvr *httpsvr.HttpServer
}

// NewService 构造方法
func NewService() (*Service, error) {

	// ka, err := kafka.NewProducer()
	// if err != nil {
	// 	return nil, err
	// }

	httpsvr := httpsvr.NewHttpServer(nil, nil)

	return &Service{
		//Kafka:   ka,
		Httpsvr: httpsvr,
	}, nil
}

func (this *Service) Start() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	//this.Kafka.Run()

	this.Httpsvr.Start()

}
