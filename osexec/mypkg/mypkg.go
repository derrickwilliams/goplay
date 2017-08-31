package mypkg

import (
	// "io/ioutil"
	"log"
	"os"
	"os/exec"
)

const (
	sourceURL = "ssh://git-codecommit.us-east-1.amazonaws.com/v1/repos/tsr-platform"
)

func Doer() {
	tmp := os.TempDir()
	log.Printf("temp %s", tmp)
	if err := os.Chdir(tmp); err != nil {
		log.Printf("couldn't CD to %s, %#v", tmp, err)
		os.Exit(-1)
		return
	}

	exec.Command("rm", "-rf", "./tsr-platform").Output()

	cmd := exec.Command("git", "clone", sourceURL)
	_, err := cmd.Output()
	if err != nil {
		log.Printf("cmd error %#v", err)
		os.Exit(-1)
		return
	}

	cmd = exec.Command("ls", "./tsr-platform")
	output, err := cmd.Output()
	if err != nil {
		log.Printf("ls error %#v", err)
		os.Exit(-1)
		return
	}

	exec.Command("rm", "-rf", "./tsr-platform").Output()

	log.Println(string(output))
}
