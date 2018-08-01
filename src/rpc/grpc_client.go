package rpc

import "sync"

type GrpcClient struct {
	sync.RWMutex
}

func NewGrpcClient() *GrpcClient {
	return &GrpcClient{}
}
