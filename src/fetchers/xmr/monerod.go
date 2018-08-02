package xmr

import (
	"github.com/dubuqingfeng/explorer-parser/src/fetchers"
	"fmt"
	"github.com/dubuqingfeng/explorer-parser/src/producer/config"
	"github.com/dubuqingfeng/explorer-parser/src/rpc"
)

type XMRFetcher struct {
	NodeConfigs []config.NodeConfig
	fetchers.Fetcher
}

func (this XMRFetcher) Fetch(title string) (bool, string) {
	fmt.Printf("test")
	// async rpc client call
	this.GetBlock("test")
	return false, "test"
}

func (this XMRFetcher) GetBlock(title string) (bool, string) {
	fmt.Println(title)
	this.RPCCall("get_block_count")
	return false, "test"
}

func (this XMRFetcher) RPCCall(method string) {
	fmt.Println(method)
	rpcClients := rpc.NewClients(this.NodeConfigs)
	rpcClients.Call(method)
}
