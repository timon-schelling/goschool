package internal

import (
	"log"
	"os/exec"
	"strings"
)

func gitRepoRoot() string {
	gitCmd := exec.Command("git", "rev-parse", "--show-toplevel")
	gitCmdOut, err := gitCmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(gitCmdOut))
}
