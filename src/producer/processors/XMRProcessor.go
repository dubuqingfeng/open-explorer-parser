package processors

import (
	"github.com/dubuqingfeng/explorer-parser/fetchers/xmr"
	"github.com/dubuqingfeng/explorer-parser/producer/config"
	"github.com/dubuqingfeng/explorer-parser/pubsub"
)

type XMRProcessor struct {
	coin   string
	status int
	reason string
	height int
	// how to rollback
	hash string
}

func NewXMRProcessor() *XMRProcessor {
	return &XMRProcessor{}
}

func (processor *XMRProcessor) Parse(message string) bool {
	// Load Fetchers
	go func() {
		// lock
		monerod := xmr.CoinXMRFetcher{NodeConfigs: config.Config.XMR.Nodes}
		// Returns an array of Object
		result, reason := monerod.Fetch("test")
		if result {
			wrapper := pubsub.NewDataWrapper("XMR", config.Config.XMR.Network, config.Config.XMR.PubConn)
			wrapper.Publish(reason)
		}
	}()
	return false
}

func (processor *XMRProcessor) Finish(info string) (status int, reason string) {
	return processor.status, processor.reason
}
