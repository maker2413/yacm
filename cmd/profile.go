package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Configuration profile creation, deletion, and management",
	Long:  `The profile command...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("profile...")
	},
}

func init() {
	rootCmd.AddCommand(profileCmd)
}
