package processors

import (
	"github.com/dubuqingfeng/explorer-parser/src/fetchers/eth"
	log "github.com/sirupsen/logrus"
	"github.com/dubuqingfeng/explorer-parser/src/producer/config"
)

type ETHProcessor struct {
	status int
	reason string
}

func NewETHProcessor() *ETHProcessor {
	return &ETHProcessor{}
}

func (this *ETHProcessor) Parse(message string) bool {
	// Load Fetchers
	log.WithField("coin_type", "ETH").Debug("Parse Start")
	go func() {
		// lock
		gethFetcher := eth.GethFetcher{NodeConfigs: config.Config.ETH.Nodes}
		// Returns an array of Object
		result, reason := gethFetcher.Fetch("test")
		log.WithField("result", result).WithField("reason", reason).Debug("test")
		// send to kafka
	}()
	return false
}

func (this *ETHProcessor) Finish(info string) (status int, reason string) {
	return this.status, this.reason
}
