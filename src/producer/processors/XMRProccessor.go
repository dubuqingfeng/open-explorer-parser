package processors

import (
	log "github.com/sirupsen/logrus"
	"github.com/dubuqingfeng/explorer-parser/src/fetchers/xmr"
)

type XMRProccessor struct {
	coinType string
	status   int
	reason   string
}

func NewXMRProcessor() *XMRProccessor {
	return &XMRProccessor{}
}

func (this *XMRProccessor) Parse(message string) bool {
	// Load Fetchers
	monerod := xmr.XMRFetcher{}
	result, reason := monerod.Fetch("test")
	log.WithField("result", result).WithField("reason", reason).Info("test")
	return false
}

func (this *XMRProccessor) Finish(info string) (status int, reason string) {
	return this.status, this.reason
}
