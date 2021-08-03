package main

import (
	"bufio"
	"fmt"
	"os"
	"portchanger/bannerusage"
	"portchanger/cli"
	"portchanger/packet"
)

func main() {
	bannerusage.Print()
	packet.Listen()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Safehouse> ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		cli.RunCommand(cmdString)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}