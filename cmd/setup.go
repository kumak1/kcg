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

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "run setup commands on each repository",
	Long:  `Running setup commands on each repository`,
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

			if err := kcgCmd.Setup(repo); err != nil {
				fmt.Println(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
	setupCmd.Flags().String("repo", "", "repository name")
	setupCmd.Flags().String("group", "", "repository group name")
	setupCmd.Flags().String("filter", "", "repository filter")
}
