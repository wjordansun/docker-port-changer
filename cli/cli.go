package cli

import (
	"os"
	"os/exec"
	"portchanger/packet"
	"strings"
)

func RunCommand(commandStr string) error {
	commandStr = strings.TrimSuffix(commandStr, "\n")
	arrCommandStr := strings.Fields(commandStr)
	switch arrCommandStr[0] {
	case "list":
		packet.OpenFile()
	case "ports":
		
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