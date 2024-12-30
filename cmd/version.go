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
	Short: "Print the version number of ledboard",
	Long:  `All software has versions. This is ledboard's'`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ledboard 0.1")
	},
}
