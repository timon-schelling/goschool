package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/timon-schelling/goschool/internal"
)

var Create = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c"},
	Short:   "create a directory based on a template",
	Long:    "create a directory based on a template located at $GOSCHOOL_HOME/templates/template_name",
	Args:    cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		template := args[0]
		name := ""
		if len(args) > 1 {
			name = args[1]
		}
		internal.Create(wd, template, name, true)
	},
}
