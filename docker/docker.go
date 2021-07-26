package docker

import (
	"fmt"
	"os/exec"
)

func Stop() {
  cmd := exec.Command("docker", "stop", "test")
  stdout, err := cmd.Output()

	if err != nil {
    fmt.Println(err.Error())
    return
  } else {
    fmt.Println(string(stdout))
  }
}

func Start() {
	cmd := exec.Command("docker", "start", "test2")
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

func Run(containerName, imageName string) {
	cmd := exec.Command("docker", "run", "-d", "--name", containerName, "-p", "8000:9999", imageName)
  stdout, err := cmd.Output()

	if err != nil {
    fmt.Println(err.Error())
    return
  } else {
    fmt.Println(string(stdout))
  }
}