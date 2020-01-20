package btc

import (
	"fmt"
	"github.com/dubuqingfeng/explorer-parser/fetchers"
	"github.com/dubuqingfeng/explorer-parser/producer/config"
	"github.com/dubuqingfeng/explorer-parser/rpc"
)

type Bitcoind struct {
	NodeConfigs []config.NodeConfig
	fetchers.Fetcher
}

func (fetcher Bitcoind) Fetch(title string) (bool, map[string]string) {
	strings := make(map[string]string)
	// async rpc client call
	_, getBlock := fetcher.GetBlock("78923")
	strings["get_block"] = getBlock
	return true, strings
}

func (fetcher Bitcoind) GetBlockCount() (bool, string) {
	fetcher.RPCCall("get_block_count", nil)
	return false, "test"
}

func (fetcher Bitcoind) GetBlock(height string) (bool, string) {
	fetcher.RPCCall("getblock", height)
	return false, "test"
}

func (fetcher Bitcoind) RPCCall(method string, param interface{}) {
	fmt.Println(method)
	rpcClients := rpc.NewClients(fetcher.NodeConfigs)
	rpcClients.Call(method, param)
}
