package xmr

import (
	"github.com/dubuqingfeng/explorer-parser/src/fetchers"
	"fmt"
	"github.com/dubuqingfeng/explorer-parser/src/producer/config"
	"github.com/dubuqingfeng/explorer-parser/src/rpc"
)

// https://getmonero.org/resources/developer-guides/daemon-rpc.html
type XMRFetcher struct {
	NodeConfigs []config.NodeConfig
	fetchers.Fetcher
}

func (this XMRFetcher) Fetch(title string) (bool, string) {
	fmt.Printf("test")
	// async rpc client call
	this.GetBlockCount()
	this.GetBlock(string(78923))
	return false, "test"
}

func (this XMRFetcher) GetBlockCount() (bool, string) {
	this.RPCCall("get_block_count", nil)
	return false, "test"
}

func (this XMRFetcher) GetBlock(height string) (bool, string) {
	this.RPCCall("getblock", height)
	return false, "test"
}

func (this XMRFetcher) RPCCall(method string, param interface{}) {
	fmt.Println(method)
	rpcClients := rpc.NewClients(this.NodeConfigs)
	rpcClients.Call(method, param)
}
