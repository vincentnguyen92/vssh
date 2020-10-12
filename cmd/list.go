package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Listing ssh server available",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing...")
	},
}
