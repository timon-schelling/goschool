package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var SaveGit = &cobra.Command{
	Use:     "git",
	Aliases: []string{"g"},
	Short:   "use git to save changes",
	Long:    "use git to save changes",
	Run: func(cmd *cobra.Command, args []string) {
		gitCmd := exec.Command("git", "rev-parse", "--show-toplevel")
		gitCmdOut, err := gitCmd.Output()
		if err != nil {
			log.Fatal(err)
		}
		gitRepoRoot := strings.TrimSpace(string(gitCmdOut))
		gitCmdAdd := exec.Command("git", "add", ".")
		gitCmdAdd.Dir = gitRepoRoot
		_ = gitCmdAdd.Run()
		gitCmdCommit := exec.Command("git", "commit", "-m", "-")
		gitCmdCommit.Dir = gitRepoRoot
		_ = gitCmdCommit.Run()
		gitCmdPush := exec.Command("git", "push")
		gitCmdPush.Dir = gitRepoRoot
		_ = gitCmdPush.Run()
		fmt.Println("Done")
	},
}
