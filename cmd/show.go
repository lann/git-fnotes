package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/lann/git-fnotes/internal"
)

var showCmd = &cobra.Command{
	Use:  "show <path>",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		obj := fmt.Sprintf(":0:%s", args[0])
		out, err := internal.Git("show", obj)
		cobra.CheckErr(err)
		fmt.Printf("%s", out)
	},
}
