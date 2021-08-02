package docker

import (
	"fmt"
	"os/exec"
	"strconv"
)

var (
	port int = 3000
	initSuccess bool = false
)

func Stop(containerName string) {
  cmd := exec.Command("docker", "stop", containerName)
  stdout, err := cmd.Output()

	if err != nil {
    fmt.Println(err.Error())
    return
  } else {
    fmt.Println(string(stdout))
  }
}

func Start(containerName string) {
	cmd := exec.Command("docker", "start", containerName)
  stdout, err := cmd.Output()

	if err != nil {
    fmt.Println(err.Error())
    return
  } else {
    fmt.Println(string(stdout))
  }
}

func Build(imageName string) {
	cmd := exec.Command("docker", "build", "-t", imageName, ".")
  stdout, err := cmd.Output()

	if err != nil {
    fmt.Println(err.Error())
    return
  } else {
    fmt.Println(string(stdout))
  }
}

func Run(containerName, imageName string, port int) {
	cmd := exec.Command("docker", "run", "-d", "--name", containerName, "-p", strconv.Itoa(port) + ":9999", imageName)
  stdout, err := cmd.Output()

	if err != nil {
    fmt.Println(err.Error())
    return
  } else {
    fmt.Println(string(stdout))
  }
}

func Init() {
	if initSuccess {
		return
	} else {
		Run("production1", "app", 3000)
		Run("honeypot1", "app", 4000)
		initSuccess = true
	}
}