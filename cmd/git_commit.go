package cmd

import (
	"github.com/spf13/cobra"

	"github.com/timon-schelling/goschool/internal"
)

var GitCommit = &cobra.Command{
	Use:     "commit",
	Aliases: []string{"c"},
	Short:   "git commit tool",
	Long:    "git commit tool",
	Run: func(cmd *cobra.Command, args []string) {
		internal.GitCommit()
	},
}
