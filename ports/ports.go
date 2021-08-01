package ports

import (
	"log"
)

var (
	dbpath string = ""
)

func ports() {

}



func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}