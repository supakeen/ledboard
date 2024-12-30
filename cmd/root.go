package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	ledcatPath string

	rootCmd = &cobra.Command{
		Use:   "ledboard",
		Short: "Show text on the `ledboard`.",
        Long:  "ledboard shows things on HUB75-compatible panels through the use of a `ledcat`-compatible binary.",
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&ledcatPath, "ledcat path", "p", "ledcat", "ledcat binary path")
}
