package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"portchanger/packet"
	"strings"
)
func runCommand(commandStr string) error {
	commandStr = strings.TrimSuffix(commandStr, "\n")
	arrCommandStr := strings.Fields(commandStr)
	switch arrCommandStr[0] {
	case "exit":
		os.Exit(0)
	case "quit":
		os.Exit(0)
	}
	cmd := exec.Command(arrCommandStr[0], arrCommandStr[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
func main() {
	packet.Listen()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Safehouse> ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		err = runCommand(cmdString)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}