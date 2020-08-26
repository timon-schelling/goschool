package cmd

import (
	"github.com/spf13/cobra"
)

var Save = &cobra.Command{
	Use:     "save",
	Aliases: []string{"s"},
	Short:   "save in a specified way",
	Long:    "save in a specified way",
}

func init() {
	Save.AddCommand(SaveGit)
}
