package app

import (
	"fmt"

	"github.com/maker2413/yacm/internal/profiles"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewInitCmd() *cobra.Command {
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Sets default profile for current system or kicks off profile creation",
		Long: `The init command will prompt the user to set a default configuration profile for
the current system. If the user doesn't have an existing profile that they want to use for the
current system init will kick off the creation of a new configuration profile.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(viper.GetString("yacm_profiles_dir"))
			p, err := profiles.Init()
			if err != nil {
				return err
			}

			files := p.GetProfiles()
			for _, p := range files {
				fmt.Println(p)
			}

			fmt.Println(p.Exist())

			return nil
		},
	}
	return initCmd
}
