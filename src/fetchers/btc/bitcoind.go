package btc

import (
	"fmt"
	"github.com/dubuqingfeng/explorer-parser/src/fetchers"
)

type Bitcoind struct {
	fetchers.Fetcher
}

//type Bitcoind interface {
//	GetAddress(address string) (err error)
//	GetTransactionByTxID(txid string)
//}

func (this Bitcoind) Fetch(title string) (bool, string) {
	fmt.Printf("test")
	// async rpc client call
	return false, "test"
}
