package app

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewBootstrapCmd() *cobra.Command {
	bootstrapCmd := &cobra.Command{
		Use:   "bootstrap",
		Short: "Kicks off system bootstrap",
		Long: `Used to bootstrap the current system with the currently
set profile or with the profile specified.

Subcommands:
  profile - Used to specify which profile to use
`,
		RunE: bootstrap,
	}

	return bootstrapCmd
}

func bootstrap(cmd *cobra.Command, args []string) error {
	fmt.Println("Bootstrap...")

	return nil
}
