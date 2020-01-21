package xmr

import (
	"github.com/dubuqingfeng/explorer-parser/fetchers"
	"github.com/dubuqingfeng/explorer-parser/producer/config"
	"github.com/dubuqingfeng/explorer-parser/rpc"
)

// https://getmonero.org/resources/developer-guides/daemon-rpc.html
type CoinXMRFetcher struct {
	NodeConfigs []config.NodeConfig
	fetchers.Fetcher
}

func (fetcher CoinXMRFetcher) Fetch(title string) (bool, map[string]string) {
	strings := make(map[string]string)
	// async rpc client call
	_, getBlockCount := fetcher.GetBlockCount()
	strings["get_block_count"] = getBlockCount
	_, getBlock := fetcher.GetBlock("18923", "")
	strings["get_block"] = getBlock
	return true, strings
}

func (fetcher CoinXMRFetcher) GetBlockCount() (bool, string) {
	result, err := fetcher.RPCCall("get_block_count", nil)
	if err != nil {
		return false, err.Error()
	}
	return true, result
}

func (fetcher CoinXMRFetcher) GetBlock(height string, hash string) (bool, string) {
	param := make(map[string]string)
	param["height"] = height
	param["hash"] = hash
	result, err := fetcher.RPCCall("getblock", param)
	if err != nil {
		return false, err.Error()
	}
	return true, result
}

func (fetcher CoinXMRFetcher) RPCCall(method string, param interface{}) (message string, err error) {
	rpcClients := rpc.NewClients(fetcher.NodeConfigs)
	message, err = rpcClients.Call(method, param)
	return message, err
}
