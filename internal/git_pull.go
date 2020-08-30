package internal

import (
	"fmt"
	"os"
	"os/exec"
)

func GitPull() {
	gitRepoRoot := gitRepoRoot()
	fmt.Println("git repo: "+gitRepoRoot)
	fmt.Println("")
	gitCmdPull := exec.Command("git", "pull")
	gitCmdPull.Dir = gitRepoRoot
	gitCmdPull.Stdout = os.Stdout
	gitCmdPull.Stderr = os.Stderr
	fmt.Println("git pull")
	_ = gitCmdPull.Run()
}
