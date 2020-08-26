package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "goschool",
	Short: "CLI tool that helps you through school and university",
}

func init() {
	rootCmd.AddCommand(Create)
	rootCmd.AddCommand(Save)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
