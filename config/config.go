package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

func init() {
	_ = os.MkdirAll(GetWorkspaceFolder(), os.ModePerm)
}

/*
GetWorkspaceFolder returns the workspace folder
~/.minecraft-manager/workspace
*/
func GetWorkspaceFolder() string {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return filepath.Join(home, folderName, "workspace")
}

/*
GetAnsibleRoleGitURL returns the Git URL to
fetch the Ansible installation role
*/
func GetAnsibleRoleGitURL() string {
	return viper.GetString("minemanager.ansible.role.git-url")
}
