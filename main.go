package main

import (
	"bufio"
	"fmt"
	"os"
	"portchanger/bannerusage"
	"portchanger/cli"
)

func main() {
	bannerusage.Print()
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