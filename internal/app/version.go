package app

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays version information for yacm",
	Long:  `The version command prints out information about yacm's current version.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("yacm version " + rootCmd.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
