package cli

import (
	"os"
	"portchanger/bannerusage"
	"portchanger/docker"
	"portchanger/packet"
	"portchanger/ports"
	"strings"
)

func RunCommand(commandStr string) {
	commandStr = strings.TrimSuffix(commandStr, "\n")
	arrCommandStr := strings.Fields(commandStr)
	switch arrCommandStr[0] {
	case "help":
		bannerusage.Print()
	case "list":
		packet.OpenFile()
	case "ports":
		ports.Display()
	case "start":
		packet.Listen()
	case "init":
		docker.Init()
	case "reset":
		docker.Reset()
	case "exit":
		os.Exit(0)
	case "quit":
		os.Exit(0)
	}
	// cmd := exec.Command(arrCommandStr[0], arrCommandStr[1:]...)
	// cmd.Stderr = os.Stderr
	// cmd.Stdout = os.Stdout
	// return cmd.Run()
}