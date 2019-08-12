package httpctx

import (
	"careyads/esserver/service/elastic"
	"careyads/esserver/service/kafka"
)

type ServerContext struct {
	EsManager *elastic.ElasticManager
	Kafka     *kafka.KafkaManager
}
