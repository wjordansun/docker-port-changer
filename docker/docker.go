package docker

import (
	"fmt"
	"os/exec"
	"strconv"
	"time"
)

var (
	//port int = 3000
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

func stopAll() {
	cmd := exec.Command("docker", "kill", "$(docker", "ps", "-q)")
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
		fmt.Println("already initialized.")
		return
	} else {
		Run("honeypot", "app", 4000)
		Run("production1", "app", 3000)

		Run("production2", "app", 3010)
		time.Sleep(2 * time.Second)
		Stop("producition2")
		time.Sleep(5 * time.Second)

		Run("production3", "app", 3020)
		time.Sleep(2 * time.Second)
		Stop("producition3")
		time.Sleep(5 * time.Second)

		Run("honeypot1", "app", 3000)
		time.Sleep(2 * time.Second)
		Stop("honeypot1")
		time.Sleep(5 * time.Second)

		Run("honeypot2", "app", 3010)
		time.Sleep(2 * time.Second)
		Stop("honeypot2")
		time.Sleep(5 * time.Second)

		Run("honeypot3", "app", 3020)
		time.Sleep(2 * time.Second)
		Stop("honeypot3")
		time.Sleep(5 * time.Second)

		fmt.Println("done initializing.")

		initSuccess = true
	}
}

func Reset() {
	stopAll()
	time.Sleep(2 * time.Second)
	Start("production1")
	Start("honeypot")
	fmt.Println("Reset complete.")
}