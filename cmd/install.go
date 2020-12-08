/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/Eldius/minecraft-manager-go/model"
	"github.com/Eldius/minecraft-manager-go/servers"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("install called")
		servers.Install(server)
	},
}

var server *model.MCServer

func init() {
	rootCmd.AddCommand(installCmd)
	server = &model.MCServer{
		Name:              "",
		Host:              "",
		User:              "",
		PrivateKeyPath:    "",
		PrivateKeyContent: "",
		ConnectionType:    model.SSHConnType,
		ConnectionPort:    "22",
	}
	installCmd.Flags().StringVarP(&server.Name, "name", "n", fmt.Sprintf("mineserver-%s", uuid.New().String()), "Server name '-n <server-name>'")
	installCmd.Flags().StringVarP(&server.User, "user", "u", "", "SSH username '-u <username>'")
	installCmd.Flags().StringVarP(&server.PrivateKeyPath, "private-key", "k", "~/.ssh/id_ed25519", "SSH Private key '-pk ~/.ssh/id_ed25519'")
	installCmd.Flags().StringVarP(&server.ConnectionPort, "port", "p", "22", "SSH port '-p 22'")

	var connType string
	installCmd.Flags().StringVarP(&connType, "connection-type", "c", "local", "Connection type '-c (SSH|LOCAL)'")

	switch strings.ToLower(connType) {
	case "ssh":
		server.ConnectionType = model.SSHConnType
	case "local":
		server.ConnectionType = model.LOCALConType
	default:
		fmt.Println("Invalid connection type")
		os.Exit(1)
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
