package docker

import (
	"fmt"
	"os/exec"
	"portchanger/badgerstuff"
	"strconv"
	"time"
)

var (
	//port int = 3000
	//initSuccess bool = false
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

func StopAll() {
	// cmd := exec.Command("docker", "kill", "$(docker", "ps", "-q)")
  // stdout, err := cmd.Output()

	// if err != nil {
  //   fmt.Println(err.Error())
  //   return
  // } else {
  //   fmt.Println(string(stdout))
  // }
	Stop("production1")
	Stop("production2")
	Stop("production3")
	Stop("honeypot1")
	Stop("honeypot2")
	Stop("honeypot3")
	//fmt.Println("Im being run.")

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
	cmd := exec.Command("docker", "run", "-d", "-it", "--name", containerName, "-p", strconv.Itoa(port) + ":9999", imageName)
  stdout, err := cmd.Output()

	if err != nil {
    fmt.Println(err.Error())
    return
  } else {
    fmt.Println(string(stdout))
  }
}

func Init() {
	if badgerstuff.DBexists(badgerstuff.DBpath) && badgerstuff.InitSuccess(){
		fmt.Println("already initialized.")
		return
	} else {
		Run("honeypot", "app", 4000)
		Stop("honeypot")

		Run("production1", "app", 3000)
		Stop("production1")

		Run("production2", "app", 3010)
		Stop("production2")

		Run("production3", "app", 3020)
		Stop("production3")

		Run("honeypot1", "app", 3000)
		Stop("honeypot1")

		Run("honeypot2", "app", 3010)
		Stop("honeypot2")

		Run("honeypot3", "app", 3020)
		Stop("honeypot3")

		Start("honeypot")
		Start("production1")

		fmt.Println("done initializing.")

		badgerstuff.Init()
	}
}

func Reset() {
	StopAll()

	//fmt.Println("Im being run.")

	time.Sleep(2 * time.Second)
	Start("production1")
	badgerstuff.SetProductionNum(1)
	fmt.Println("Reset complete.")

}