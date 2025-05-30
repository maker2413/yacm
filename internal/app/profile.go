package app

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewProfileCmd() *cobra.Command {
	profileCmd := &cobra.Command{
		Use:   "profile",
		Short: "Configuration profile creation, deletion, and management",
		Long: `Used to create and manage yacm profiles

Subcommands:
  create  - Creates a new profile
  default - Sets default profile for current system
  delete  - Deletes a profile
  list    - List profiles on system
  help    - Prints this menu

Alternative subcommands:
  --help  - Prints this menu
  --rm    - Deletes a profile
  -h      - Prints this menu
  -l      - List profiles on system
`,
		RunE: executeProfile,
	}

	return profileCmd
}

func executeProfile(cmd *cobra.Command, args []string) error {
	fmt.Println("profile...")

	return nil
}
