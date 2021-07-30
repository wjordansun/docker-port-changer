package ipfs

import (
	"fmt"
	"os/exec"
)

func Add() {
	cmd := exec.Command("ipfs", "add", "test.pcap")
  stdout, err := cmd.Output()

	if err != nil {
    fmt.Println(err.Error())
    return
  } else {
    fmt.Println(string(stdout))
  }
}

func Cat() {
	
}