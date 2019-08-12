package httpctx

import (
	"logcat/consumer/service/elastic"
)

type ServerContext struct {
	EsManager *elastic.ElasticManager
}
