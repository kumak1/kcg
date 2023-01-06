/*
Copyright © 2023 kumak1

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
	"github.com/kumak1/kcg/kcg"
	kcgExec "github.com/kumak1/kcg/kcg/exec"

	"github.com/spf13/cobra"
)

// cleanupCmd represents the cleanup command
var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "delete merged branch on each repository dir",
	Long:  `Delete merged branch on each repository dir`,
	Run: func(cmd *cobra.Command, args []string) {
		repoFlag, _ := cmd.Flags().GetString("repo")
		groupFlag, _ := cmd.Flags().GetString("group")
		filterFlag, _ := cmd.Flags().GetString("filter")
		kcgCmd := kcg.Command(config)

		for index, repo := range kcgCmd.List(repoFlag, groupFlag, filterFlag) {
			if path := kcgCmd.Path(repo); !kcgExec.DirExists(path) {
				fmt.Println("invalid path: '" + index + "' " + path)
				continue
			}

			if err := kcgCmd.Cleanup(repo); err != nil {
				fmt.Println(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(cleanupCmd)
	cleanupCmd.Flags().String("repo", "", "repository name")
	cleanupCmd.Flags().String("group", "", "repository group name")
	cleanupCmd.Flags().String("filter", "", "repository filter")
}
