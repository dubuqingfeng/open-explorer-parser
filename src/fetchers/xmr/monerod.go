package xmr

import (
	"github.com/dubuqingfeng/explorer-parser/src/fetchers"
	"fmt"
)

type XMRFetcher struct {
	fetchers.Fetcher
}

func (this XMRFetcher) Fetch(title string) (bool, string) {
	fmt.Printf("test")
	// async rpc client call
	return false, "test"
}
