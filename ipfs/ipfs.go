package ipfs

import (
	"fmt"
	"os/exec"
)

func Add(fileName string) string{
	cmd := exec.Command("ipfs", "add", fileName)
  stdout, err := cmd.Output()

	if err != nil {
    fmt.Println(err.Error())
    return err.Error()
  } else {
    return string(stdout)
  }
}

func Cat(CID string) string{
	cmd := exec.Command("ipfs", "cat", CID, ">", "sample.pcap")
  stdout, err := cmd.Output()

	if err != nil {
    fmt.Println(err.Error())
    return err.Error()
  } else {
    return string(stdout)
  }
}

// func ipfsToPcap(output string) {
// 	cmd := exec.Command("ipfs", "cat", CID)
//   stdout, err := cmd.Output()

// 	if err != nil {
//     fmt.Println(err.Error())
//     return err.Error()
//   } else {
//     return string(stdout)
//   }
// }