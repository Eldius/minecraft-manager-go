package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const (
	folderName = ".minecraft-manager"
)

/*
SetupViper sets up the Viper configuration
*/
func SetupViper(cfgFile string) {
	SetDefaults()
	BindEnvVars()
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(filepath.Join(home, folderName))
		viper.AddConfigPath(".")
		viper.SetConfigName("minecraft-manager")
		viper.SetConfigType("yml")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

}

/*
SetDefaults set's up configuration default values
*/
func SetDefaults() {
	viper.SetDefault(
		"minemanager.ansible.role.git-url",
		"https://github.com/Eldius/minecraft-java-edition-ansible-role.git",
	)
	viper.SetDefault("minemanager.mojang.versions-url", "https://launchermeta.mojang.com/mc/game/version_manifest.json")
}

/*
BindEnvVars binds configuration keys to env vars
*/
func BindEnvVars() {

}
