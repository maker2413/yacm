package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Kicks off system bootstrap",
	Long: `Used to bootstrap the current system with the currently
set profile or with the profile specified.

Subcommands:
  profile - Used to specify which profile to use
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Bootstrap...")
	},
}

func init() {
	rootCmd.AddCommand(bootstrapCmd)
}
