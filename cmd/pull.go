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

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "run `git pull` on each repository dir",
	Long:  `Running git pull command on each repository dir`,
	Run: func(cmd *cobra.Command, args []string) {
		repoFlag, _ := cmd.Flags().GetString("repo")
		groupFlag, _ := cmd.Flags().GetString("group")
		gitCommand := kcg.GitCommand(config)

		for index, repo := range gitCommand.List(repoFlag, groupFlag, "") {
			if path := gitCommand.Path(repo); !kcgExec.DirExists(path) {
				fmt.Println("invalid path: '" + index + "' " + path)
				continue
			}

			if err := gitCommand.Pull(repo); err != nil {
				fmt.Println(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
	pullCmd.Flags().String("repo", "", "repository name")
	pullCmd.Flags().String("group", "", "repository group name")
}
