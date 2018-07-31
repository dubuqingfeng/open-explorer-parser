package rpc

import "net/http"

type rpcClient struct {
	address    string
	user       string
	passwd     string
	ssl        bool
	httpClient *http.Client
}

type rpcResponse struct {
	message string
}

type rpcRequest struct {
}
