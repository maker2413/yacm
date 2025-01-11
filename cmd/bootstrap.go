package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Kicks off system bootstrap",
	Long:  `The bootstrap command...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Bootstrap...")
	},
}
