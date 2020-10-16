package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "vssh",
		Short: "Vincent SSH Connections",
		Long:  `A tools for management ssh connections.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
