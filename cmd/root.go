package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "git-fnotes"}

func init() {
  rootCmd.AddCommand(showCmd)
}

func Execute() {
  rootCmd.Execute()
}
