package filters

import (
	"github.com/dubuqingfeng/explorer-parser/src/pubsub"
	"github.com/dubuqingfeng/explorer-parser/src/consumer/config"
)

type BTCFilter struct {
	status int
	reason string
}

func NewBTCFilter() *BTCFilter {
	return &BTCFilter{}
}

func (filter *BTCFilter) Filter(message string) bool {
	// Load Filter
	for {
		go func() {
			wrapper := pubsub.NewDataWrapper("XMR", config.Config.XMR.Network, config.Config.XMR.PubConn)
			wrapper.Subscribe("channel")
		}()
	}
	return false
}
