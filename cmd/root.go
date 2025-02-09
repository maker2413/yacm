package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "yacm",
	Short: "yacm (Yet Another Configuration Manager)",
	Long: `yacm is a CLI tool that allows users to configure and manage all of their
computers with simple yaml manifests.`,
}

func Execute(version string) error {
	rootCmd.Version = version
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/yacm/config.yml)")
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	homeDir, err := os.UserHomeDir()
	cobra.CheckErr(err)

	viper.SetDefault("yacm_base_dir", fmt.Sprintf("%s/.config", homeDir))
	viper.SetDefault("yacm_dir", fmt.Sprintf("%s/yacm", viper.GetString("yacm_base_dir")))
	viper.SetDefault("yacm_profiles_dir", fmt.Sprintf("%s/profiles", viper.GetString("yacm_dir")))
	viper.SetDefault("yacm_scripts_dir", fmt.Sprintf("%s/scripts", viper.GetString("yacm_dir")))

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
		err := viper.ReadInConfig()
		cobra.CheckErr(err)
	} else {
		defaultConfigPath := path.Join(viper.GetString("yacm_dir"), "config.yml")
		viper.SetConfigFile(defaultConfigPath)

		// if we error trying to read the config file it must need to be generated
		if err := viper.ReadInConfig(); err != nil {
			viper.SafeWriteConfigAs(defaultConfigPath)
			err = viper.ReadInConfig()
			cobra.CheckErr(err)
		}
	}

	// read in environment variables that match
	viper.AutomaticEnv()
}
