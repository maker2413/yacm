package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Sets default profile for current system or kicks off profile creation",
	Long: `The init command will prompt the user to set a default configuration profile for
the current system. If the user doesn't have an existing profile that they want to use for the
current system init will kick off the creation of a new configuration profile.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init...")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
