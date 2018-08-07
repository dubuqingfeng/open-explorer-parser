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

func (fetcher GethFetcher) Fetch(title string) (bool, map[string]string) {
	strings := make(map[string]string)
	// async rpc client call
	_, getBlock := fetcher.GetBlock("78923")
	strings["get_block"] = getBlock
	return true, strings
}

func (fetcher GethFetcher) GetBlock(title string) (bool, string) {
	fmt.Println(title)
	fetcher.RPCCall("rpc client call method", "")
	return false, "test"
}

func (fetcher GethFetcher) RPCCall(method string, param interface{}) {
	fmt.Println(method)
	rpcClients := rpc.NewClients(fetcher.NodeConfigs)
	rpcClients.Call(method, param)
}
