package app

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
func NewVersionCmd() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Displays version information for yacm",
		Long:  `The version command prints out information about yacm's current version.`,
		RunE:  printVersion,
	}

	return versionCmd
}

func printVersion(cmd *cobra.Command, args []string) error {
	if cmd.Parent() == nil {
		return errors.New("Unable to read Parent()")
	}

	if cmd.Parent().Version == "" {
		return errors.New("Unable to read Parent().Version")
	}

	fmt.Printf("yacm version %s\n", cmd.Parent().Version)

	return nil
}
