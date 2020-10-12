package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version v0.1",
	Long:  `First version v0.1`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You are using the firstly version. v0.1")
	},
}
