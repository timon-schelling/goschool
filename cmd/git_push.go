package cmd

import (
	"github.com/spf13/cobra"

	"github.com/timon-schelling/goschool/internal"
)

var GitPush = &cobra.Command{
	Use:     "push",
	Aliases: []string{"ps"},
	Short:   "git push tool",
	Long:    "git push tool",
	Run: func(cmd *cobra.Command, args []string) {
		internal.GitPush()
	},
}
