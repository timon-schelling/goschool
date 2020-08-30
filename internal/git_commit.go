package internal

import (
	"fmt"
	"os"
	"os/exec"
)

func GitCommit() {
	gitRepoRoot := gitRepoRoot()
	fmt.Println("git repo: "+gitRepoRoot)
	fmt.Println("")
	gitCmdCommit := exec.Command("git", "commit", "-m", "-")
	gitCmdCommit.Dir = gitRepoRoot
	gitCmdCommit.Stdout = os.Stdout
	gitCmdCommit.Stderr = os.Stderr
	fmt.Println("git commit")
	_ = gitCmdCommit.Run()
}
