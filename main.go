package main

import (
	"bufio"
	"fmt"
	"os"
	"portchanger/cli"
)

func main() {
	//packet.Listen()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Safehouse> ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		err = cli.RunCommand(cmdString)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}