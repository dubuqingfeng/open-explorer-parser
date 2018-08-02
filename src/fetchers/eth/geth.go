package eth

import (
	"github.com/dubuqingfeng/explorer-parser/src/fetchers"
	"fmt"
	"github.com/dubuqingfeng/explorer-parser/src/rpc"
	"github.com/dubuqingfeng/explorer-parser/src/producer/config"
)

type GethFetcher struct {
	// Node config
	NodeConfigs []config.NodeConfig
	fetchers.Fetcher
}

func (this GethFetcher) Fetch(title string) (bool, string) {
	fmt.Println("fetch")
	this.GetBlock("get block")
	return false, "test"
}

func (this GethFetcher) GetBlock(title string) (bool, string) {
	fmt.Println(title)
	this.RPCCall("rpc client call method")
	return false, "test"
}

func (this GethFetcher) RPCCall(method string) {
	fmt.Println(method)
	rpcClients := rpc.NewClients(this.NodeConfigs)
	rpcClients.Call(method)
}
