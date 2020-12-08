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

	"github.com/Eldius/minecraft-manager-go/versions"
	"github.com/spf13/cobra"
)

// versionsCmd represents the versions command
var versionsCmd = &cobra.Command{
	Use:   "versions",
	Short: "List available versions",
	Long:  `List available versions.`,
	Run: func(cmd *cobra.Command, args []string) {
		if vers, err := versions.GetVersions(); err != nil {
			fmt.Println("Failed to list available versions")
		} else {
			fmt.Printf("## available versions: (%d)\n", len(vers))
			for _, v := range vers {
				fmt.Printf("  - %s from (%s)\n", v.ID, v.ReleaseTime)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(versionsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
