package rpc

func NewClients() []*RpcClient {
	// Need to support custom clients
	clients := make([]*RpcClient, 0)
	rpc := newRPCClient("address", "user", "passwd", false)
	clients = append(clients, rpc)
	return clients
}

func (this *RpcClient) Call(rpc RpcClient) (message string, err error) {
	// waitgroup
	return "test", nil
}
