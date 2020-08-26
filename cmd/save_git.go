package cmd

import (
	"github.com/spf13/cobra"

	"github.com/timon-schelling/goschool/internal"
)

var SaveGit = &cobra.Command{
	Use:     "git",
	Aliases: []string{"g"},
	Short:   "use git to save changes",
	Long:    "use git to save changes",
	Run: func(cmd *cobra.Command, args []string) {
		internal.SaveGit()
	},
}
