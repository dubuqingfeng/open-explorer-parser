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
	go func() {
		wrapper := pubsub.NewDataWrapper("XMR", config.Config.XMR.Network, config.Config.XMR.PublishType, config.Config.XMR.PubConn)
		wrapper.Subscribe("channel")
		//pubsub := client.Subscribe(channel)
		//msg,_ := pubsub.Receive()
		//fmt.Println("Receive from channel:", msg)
		//done <- struct {}{}
	}()
	return false
}
