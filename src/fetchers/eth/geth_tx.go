package eth

import (
	"github.com/dubuqingfeng/explorer-parser/src/fetchers"
	"fmt"
	"github.com/dubuqingfeng/explorer-parser/src/producer/config"
)

type GethTxFetcher struct {
	NodeConfigs []config.NodeConfig
	fetchers.Fetcher
}

func (this GethTxFetcher) Fetch(title string) (bool, string) {
	fmt.Println("fetch")
	// async rpc client call
	this.RPCCall("rpc client call method")
	return false, "test"
}
