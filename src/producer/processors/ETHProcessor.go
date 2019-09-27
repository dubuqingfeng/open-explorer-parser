package processors

import (
	"github.com/dubuqingfeng/explorer-parser/src/fetchers/eth"
	"github.com/dubuqingfeng/explorer-parser/src/producer/config"
	"github.com/dubuqingfeng/explorer-parser/src/pubsub"
	log "github.com/sirupsen/logrus"
)

type ETHProcessor struct {
	status int
	reason string
}

func NewETHProcessor() *ETHProcessor {
	return &ETHProcessor{}
}

func (processor *ETHProcessor) Parse(message string) bool {
	// Load Fetchers
	log.WithField("coin_type", "ETH").Debug("Parse Start")
	go func() {
		// lock
		gethFetcher := eth.GethFetcher{NodeConfigs: config.Config.ETH.Nodes}
		// Returns an array of Object
		result, reason := gethFetcher.Fetch("test")
		log.WithField("result", result).WithField("reason", reason).Debug("test")
		// send to kafka
		wrapper := pubsub.NewDataWrapper("ETH", config.Config.ETH.Network, config.Config.ETH.PubConn)
		wrapper.Publish(reason)
	}()
	return false
}

func (processor *ETHProcessor) Finish(info string) (status int, reason string) {
	return processor.status, processor.reason
}
