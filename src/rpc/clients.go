package rpc

import (
	"fmt"
)

type rpcClients struct {
	clients []*RpcClient
}

func NewClients() *rpcClients {
	// Need to support custom clients
	clients := make([]*RpcClient, 0)
	rpc := newRPCClient("address", "user", "passwd", false)
	clients = append(clients, rpc)
	return &rpcClients{clients: clients}
}

func (this *rpcClients) Call(method string) (message string, err error) {
	fmt.Println("rpc clients call")
	for _, value := range this.clients {
		// avoid goroutine leak
		go func(client *RpcClient, method string) {
			client.call(method, "param")
		}(value, method)
	}
	return "test", nil
}
