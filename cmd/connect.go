package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(connectCmd)
}

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to a SSH server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Connection...")
	},
}
