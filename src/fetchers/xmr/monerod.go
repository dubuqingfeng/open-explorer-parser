package xmr

import (
	"github.com/dubuqingfeng/explorer-parser/src/fetchers"
	"github.com/dubuqingfeng/explorer-parser/src/producer/config"
	"github.com/dubuqingfeng/explorer-parser/src/rpc"
)

// https://getmonero.org/resources/developer-guides/daemon-rpc.html
type XMRFetcher struct {
	NodeConfigs []config.NodeConfig
	fetchers.Fetcher
}

func (this XMRFetcher) Fetch(title string) (bool, map[string]string) {
	strings := make(map[string]string)
	// async rpc client call
	_, get_block_count := this.GetBlockCount()
	strings["get_block_count"] = get_block_count
	_, get_block := this.GetBlock("78923", "")
	strings["get_block"] = get_block
	return true, strings
}

func (this XMRFetcher) GetBlockCount() (bool, string) {
	result, err := this.RPCCall("get_block_count", nil)
	if err != nil {
		return false, err.Error()
	}
	return true, result
}

func (this XMRFetcher) GetBlock(height string, hash string) (bool, string) {
	param := make(map[string]string)
	param["height"] = height
	param["hash"] = hash
	result, err := this.RPCCall("getblock", param)
	if err != nil {
		return false, err.Error()
	}
	return true, result
}

func (this XMRFetcher) RPCCall(method string, param interface{}) (message string, err error) {
	rpcClients := rpc.NewClients(this.NodeConfigs)
	message, err = rpcClients.Call(method, param)
	return message, err
}
