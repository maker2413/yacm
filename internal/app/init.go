package app

import (
	"fmt"

	"github.com/maker2413/yacm/internal/profiles"
	"github.com/spf13/cobra"
)

func NewInitCmd() *cobra.Command {
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Sets default profile for current system or kicks off profile creation",
		Long: `The init command will prompt the user to set a default configuration profile for
the current system. If the user doesn't have an existing profile that they want to use for the
current system init will kick off the creation of a new configuration profile.`,
		RunE: executeInit,
	}
	return initCmd
}

func executeInit(cmd *cobra.Command, args []string) error {
	p, err := profiles.Init()
	if err != nil {
		return err
	}

	for i, profile := range p.GetProfiles() {
		fmt.Println(i, profile)
	}

	fmt.Println(p.Exist())

	return nil
}
