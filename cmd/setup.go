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

	"github.com/spf13/cobra"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "run configfile setupCommands each repository",
	Long:  `Running configfile setupCommands each repository`,
	Run: func(cmd *cobra.Command, args []string) {
		repoFlag, _ := cmd.Flags().GetString("repo")
		gitCommand := kcg.GitCommand(config)

		for index, repo := range config.Repos {
			if repoFlag != "" && repoFlag != index {
				continue
			}

			err := gitCommand.Setup(repo)
			if err != nil {
				fmt.Println(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
	setupCmd.Flags().String("repo", "", "repository name")
}
