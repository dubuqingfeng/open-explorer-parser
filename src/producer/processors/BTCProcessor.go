package processors

import (
	log "github.com/sirupsen/logrus"
	"github.com/dubuqingfeng/explorer-parser/src/fetchers/btc"
	"github.com/dubuqingfeng/explorer-parser/src/producer/config"
)

type BTCProcessor struct {
	status int
	reason string
}

func NewBTCProcessor() *BTCProcessor {
	return &BTCProcessor{}
}

func (this *BTCProcessor) Parse(message string) bool {
	// Load Fetchers
	log.WithField("coin_type", "BTC").Debug("Parse Start")
	go func() {
		// lock
		monerod := btc.Bitcoind{NodeConfigs: config.Config.BTC.Nodes}
		// Returns an array of Object
		result, reason := monerod.Fetch("test")
		log.WithField("result", result).WithField("reason", reason).Debug("test")
		// send to kafka
	}()
	return false
}

func (this *BTCProcessor) Finish(info string) (status int, reason string) {
	return this.status, this.reason
}
