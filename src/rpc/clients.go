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
		rpc := newRPCClient(value)
		clients = append(clients, rpc)
	}
	return &rpcClients{clients: clients}
}

func (this *rpcClients) Call(method string, param interface{}) (message string, err error) {
	fmt.Println("rpc clients call")
	for _, value := range this.clients {
		result, err := value.call(method, param)
		if err == nil {
			return result.JSONRPC, nil
		}
	}
	return "test", nil
}
