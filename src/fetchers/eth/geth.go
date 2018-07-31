package eth

import (
	"github.com/dubuqingfeng/explorer-parser/src/fetchers"
	"fmt"
)

type GethFetcher struct {
	fetchers.Fetcher
}

func (this GethFetcher) Fetch(title string) (bool, string) {
	fmt.Printf("test")
	return false, "test"
}
