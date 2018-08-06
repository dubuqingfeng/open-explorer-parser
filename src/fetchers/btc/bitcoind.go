package btc

import (
	"fmt"
	"github.com/dubuqingfeng/explorer-parser/src/fetchers"
	"github.com/dubuqingfeng/explorer-parser/src/rpc"
	"github.com/dubuqingfeng/explorer-parser/src/producer/config"
)

type Bitcoind struct {
	NodeConfigs []config.NodeConfig
	fetchers.Fetcher
}

func (this Bitcoind) Fetch(title string) (bool, string) {
	fmt.Printf("test")
	// async rpc client call
	this.GetBlockCount()
	this.GetBlock(string(78923))
	return false, "test"
}

func (this Bitcoind) GetBlockCount() (bool, string) {
	this.RPCCall("get_block_count", nil)
	return false, "test"
}

func (this Bitcoind) GetBlock(height string) (bool, string) {
	this.RPCCall("getblock", height)
	return false, "test"
}

func (this Bitcoind) RPCCall(method string, param interface{}) {
	fmt.Println(method)
	rpcClients := rpc.NewClients(this.NodeConfigs)
	rpcClients.Call(method, param)
}
