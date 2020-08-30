package cmd

import (
	"github.com/spf13/cobra"

	"github.com/timon-schelling/goschool/internal"
)

var GitPull = &cobra.Command{
	Use:     "pull",
	Aliases: []string{"pl"},
	Short:   "git pull tool",
	Long:    "git pull tool",
	Run: func(cmd *cobra.Command, args []string) {
		internal.GitPull()
	},
}
