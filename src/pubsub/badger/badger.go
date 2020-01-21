package badger

import (
	"github.com/dgraph-io/badger"
	log "github.com/sirupsen/logrus"
)

func InitBadger() {
	db, err := badger.Open(badger.DefaultOptions("/tmp/badger"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
