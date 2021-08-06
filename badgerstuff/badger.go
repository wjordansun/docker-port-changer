package badgerstuff

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/dgraph-io/badger"
)

const (
	DBpath = "./tmp/data"
)

var (
	value []byte
	productionNum []byte
	prodNum int
)

func DBexists(path string) bool {
	if _, err := os.Stat(path + "/MANIFEST"); os.IsNotExist(err) {
		return false
	}

	return true
}

func Init() {
	opts := badger.DefaultOptions(DBpath)
	db, err := badger.Open(opts)
	Handle(err)
	defer db.Close()
	
	// txn := db.NewTransaction(true) // Read-write txn
	// err = txn.SetEntry(badger.NewEntry([]byte("initSuccess"), []byte("true")))
	// Handle(err)

	// err = txn.Commit()
	// Handle(err)

	err = db.Update(func(txn *badger.Txn) error {
		if _, err := txn.Get([]byte("initSuccess")); err == badger.ErrKeyNotFound { //database has not been created yet
			err = txn.Set([]byte("initSuccess"), []byte("true"))
			fmt.Println("database created.")
			Handle(err)

			err = txn.Set([]byte("ProductionNum"), []byte("1"))
			return err
		}
		return err
	})
	Handle(err)
}

func InitSuccess() bool {
	opts := badger.DefaultOptions(DBpath)
	db, err := badger.Open(opts)
	Handle(err)
	defer db.Close()

	// err = db.View(func(txn *badger.Txn) error {
	// 	item, err := txn.Get([]byte("initSuccess"))
	// 	Handle(err)
	// 	val, err = item.ValueCopy(nil)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	return nil
	// })
	// Handle(err)

	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("initSuccess"))
		Handle(err)
		err = item.Value(func(val []byte) error {
			value = append([]byte{}, val...)
			return err
		})

		return nil
	})
	Handle(err)
	

	if string(value) == "true" {
		return true
	} else {
		return false
	}

}

func ProductionNum() int {
	opts := badger.DefaultOptions(DBpath)
	db, err := badger.Open(opts)
	Handle(err)
	defer db.Close()

	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("ProductionNum"))
		Handle(err)
		err = item.Value(func(val []byte) error {
			productionNum = append([]byte{}, val...)
			return err
		})

		return nil
	})
	Handle(err)

	switch string(productionNum) {
	case "1":
		prodNum = 1
	case "2":
		prodNum = 2
	case "3":
		prodNum = 3
	}

	return prodNum

}

func SetProductionNum(num int) {
	opts := badger.DefaultOptions(DBpath)
	db, err := badger.Open(opts)
	Handle(err)
	defer db.Close()

	err = db.Update(func(txn *badger.Txn) error {
		err := txn.Delete([]byte("ProductionNum"))
		Handle(err)
		
		err = txn.Set([]byte("ProductionNum"), []byte(strconv.Itoa(num)))
		return err
	})
	Handle(err)

}

func Handle(err error) {
	if err !=nil {
		log.Panic(err)
	}
}