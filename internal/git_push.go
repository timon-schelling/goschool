package internal

import (
	"fmt"
	"os"
	"os/exec"
)

func GitPush() {
	gitRepoRoot := gitRepoRoot()
	fmt.Println("git repo: "+gitRepoRoot)
	fmt.Println("")
	gitCmdPush := exec.Command("git", "push")
	gitCmdPush.Dir = gitRepoRoot
	gitCmdPush.Stdout = os.Stdout
	gitCmdPush.Stderr = os.Stderr
	fmt.Println("git push")
	_ = gitCmdPush.Run()
}
