package processors

import (
	log "github.com/sirupsen/logrus"
	"github.com/dubuqingfeng/explorer-parser/src/fetchers/xmr"
	"github.com/dubuqingfeng/explorer-parser/src/producer/config"
)

type XMRProcessor struct {
	coinType string
	status   int
	reason   string
}

func NewXMRProcessor() *XMRProcessor {
	return &XMRProcessor{}
}

func (this *XMRProcessor) Parse(message string) bool {
	// Load Fetchers
	log.WithField("coin_type", "XMR").Debug("Parse Start")
	go func() {
		// lock
		monerod := xmr.XMRFetcher{NodeConfigs: config.Config.XMR.Nodes}
		// Returns an array of Object
		result, reason := monerod.Fetch("test")
		log.WithField("result", result).WithField("reason", reason).Debug("test")
		// send to kafka
	}()
	return false
}

func (this *XMRProcessor) Finish(info string) (status int, reason string) {
	return this.status, this.reason
}
