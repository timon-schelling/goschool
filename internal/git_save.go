package internal

import (
	"fmt"
	"os"
	"os/exec"
)

func GitSave()  {
	gitRepoRoot := gitRepoRoot()
	fmt.Println("git repo: "+gitRepoRoot)
	fmt.Println("")

	gitCmdPull := exec.Command("git", "pull")
	gitCmdPull.Dir = gitRepoRoot
	gitCmdPull.Stdout = os.Stdout
	gitCmdPull.Stderr = os.Stderr
	fmt.Println("git pull")
	_ = gitCmdPull.Run()
	fmt.Println()

	gitCmdAdd := exec.Command("git", "add", ".")
	gitCmdAdd.Dir = gitRepoRoot
	gitCmdAdd.Stdout = os.Stdout
	gitCmdAdd.Stderr = os.Stderr
	fmt.Println("git add")
	_ = gitCmdAdd.Run()
	fmt.Println()

	gitCmdCommit := exec.Command("git", "commit", "-m", "-")
	gitCmdCommit.Dir = gitRepoRoot
	gitCmdCommit.Stdout = os.Stdout
	gitCmdCommit.Stderr = os.Stderr
	fmt.Println("git commit")
	_ = gitCmdCommit.Run()
	fmt.Println()

	gitCmdPush := exec.Command("git", "push")
	gitCmdPush.Dir = gitRepoRoot
	gitCmdPush.Stdout = os.Stdout
	gitCmdPush.Stderr = os.Stderr
	fmt.Println("git push")
	_ = gitCmdPush.Run()
}
