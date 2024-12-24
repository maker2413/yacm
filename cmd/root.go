package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "yacm",
	Version: yacmVersion,
	Short:   "yacm (Yet Another Configuration Manager)",
	Long: `yacm is a CLI tool that allows users to configure and manage all of their
computers with simple yaml manifests.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
