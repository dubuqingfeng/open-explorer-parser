package rpc

import (
	"fmt"
	"github.com/dubuqingfeng/explorer-parser/src/producer/config"
)

type rpcClients struct {
	clients []*RpcClient
}

func NewClients(nodeConfigs []config.NodeConfig) *rpcClients {
	// Need to support custom clients
	clients := make([]*RpcClient, 0)
	for _, value := range nodeConfigs {
		rpc := newRPCClient(value.Address, value.User, value.Password, value.SSL)
		clients = append(clients, rpc)
	}
	return &rpcClients{clients: clients}
}

func (this *rpcClients) Call(method string) (message string, err error) {
	fmt.Println("rpc clients call")
	for _, value := range this.clients {
		result, err := value.call(method, "param")
		if err == nil {
			return result.JSONRPC, nil
		}
		// avoid goroutine leak
		//go func(client *RpcClient, method string) {
		//}(value, method)
	}
	return "test", nil
}
