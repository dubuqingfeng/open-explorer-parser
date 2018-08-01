package eth

import (
	"github.com/dubuqingfeng/explorer-parser/src/fetchers"
	"fmt"
	"github.com/dubuqingfeng/explorer-parser/src/rpc"
)

type GethFetcher struct {
	fetchers.Fetcher
}

func (this GethFetcher) Fetch(title string) (bool, string) {
	fmt.Println("fetch")
	// async rpc client call
	go rpcCall("rpc client call method")
	return false, "test"
}

func rpcCall(method string) {
	fmt.Println(method)
	rpcClients := rpc.NewClients()
	rpcClients.Call(method)
}
