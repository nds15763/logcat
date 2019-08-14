package main

import (
	"beats/src/crawler"
	"beats/src/publisher"
)

func main() {

	crawler, err := crawler.NewCrawler(config.gCfg)
	if err != nil {
		return
	}

	publisher, err := publisher.NewPublisher(crawler)
	if err != nil {
		return
	}
	publisher.Run()

}
