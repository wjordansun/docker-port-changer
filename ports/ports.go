package ports

import (
	"fmt"
	"log"
	"os/exec"
)

var (
	dbpath string = ""
)

func Display() {
	// switch packet.ProductionNum {
	// case 1:
	// 	fmt.Println("Honeypot server on port 4000")
	// 	fmt.Println("Honeypot server on port 3020")
	// 	fmt.Println("Production server on port 3000")
	// case 2:
	// 	fmt.Println("Honeypot server on port 4000")
	// 	fmt.Println("Honeypot server on port 3000")
	// 	fmt.Println("Production server on port 3010")
	// case 3:
	// 	fmt.Println("Honeypot server on port 4000")
	// 	fmt.Println("Honeypot server on port 3010")
	// 	fmt.Println("Production server on port 3020")
	// }

	cmd := exec.Command("docker", "ps", "-a")
  stdout, err := cmd.Output()

	if err != nil {
    fmt.Println(err.Error())
    return
  } else {
    fmt.Println(string(stdout))
  }

}



func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}