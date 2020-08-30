package cmd

import (
	"github.com/spf13/cobra"
)

var Git = &cobra.Command{
	Use:     "git",
	Aliases: []string{"g"},
	Short:   "git tools",
	Long:    "git tools",
}

func init() {
	Git.AddCommand(GitAdd)
	Git.AddCommand(GitCommit)
	Git.AddCommand(GitPush)
	Git.AddCommand(GitPull)
	Git.AddCommand(GitSave)
}
