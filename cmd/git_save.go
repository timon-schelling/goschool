package cmd

import (
	"github.com/spf13/cobra"

	"github.com/timon-schelling/goschool/internal"
)

var GitSave = &cobra.Command{
	Use:     "save",
	Aliases: []string{"s"},
	Short:   "git save tool",
	Long:    "git save tool",
	Run: func(cmd *cobra.Command, args []string) {
		internal.GitSave()
	},
}
