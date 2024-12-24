package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var yacmVersion = "v0.3.0"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays version information for yacm",
	Long:  `The version command prints out information about yacm's current version.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("yacm version " + yacmVersion)
	},
}
