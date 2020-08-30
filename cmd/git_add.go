package cmd

import (
	"github.com/spf13/cobra"

	"github.com/timon-schelling/goschool/internal"
)

var all bool

var GitAdd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "git add tool",
	Long:    "git add tool",
	Run: func(cmd *cobra.Command, args []string) {
		if all {
			internal.GitAddAll()
		} else {
			internal.GitAdd()
		}
	},
}

func init() {
	GitAdd.Flags().BoolVarP(&all, "all", "a", false, "add all from repo root")
}
