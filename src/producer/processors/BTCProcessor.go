package processors

import (
	log "github.com/sirupsen/logrus"
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
	return false
}

func (this *BTCProcessor) Finish(info string) (status int, reason string) {
	return this.status, this.reason
}
