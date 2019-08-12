package httpctx

import (
	"logcat/logcatch/service/elastic"
	"logcat/logcatch/service/kafka"
	"logcat/tools/logrus"

	"github.com/kataras/iris"
)

type IrisRequest struct {
	Ctx     iris.Context
	Logger  *logrus.Logger
	Context *ServerContext
}

type ServerContext struct {
	EsManager *elastic.ElasticManager
	Kafka     *kafka.Producer
}
