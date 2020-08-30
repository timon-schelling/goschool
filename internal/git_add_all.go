package internal

import (
	"fmt"
	"os"
	"os/exec"
)

func GitAddAll() {
	gitRepoRoot := gitRepoRoot()
	fmt.Println("git repo: "+gitRepoRoot)
	fmt.Println("")
	gitCmdAdd := exec.Command("git", "add", ".")
	gitCmdAdd.Dir = gitRepoRoot
	gitCmdAdd.Stdout = os.Stdout
	gitCmdAdd.Stderr = os.Stderr
	fmt.Println("git add")
	_ = gitCmdAdd.Run()
}
