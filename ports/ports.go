package ports

import (
	"fmt"
	"log"
	"portchanger/packet"
)

var (
	dbpath string = ""
)

func Display() {
	switch packet.ProductionNum {
	case 1:
		fmt.Println("Honeypot server on port 4000")
		fmt.Println("Honeypot server on port 3020")
		fmt.Println("Production server on port 3000")
	case 2:
		fmt.Println("Honeypot server on port 4000")
		fmt.Println("Honeypot server on port 3000")
		fmt.Println("Production server on port 3010")
	case 3:
		fmt.Println("Honeypot server on port 4000")
		fmt.Println("Honeypot server on port 3010")
		fmt.Println("Production server on port 3020")
	}
}



func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}