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

// switchCmd represents the switch command
var switchCmd = &cobra.Command{
	Use:   "switch <branch>",
	Short: "run `git switch` on each repository dir",
	Long:  `Running git switch command on each repository dir`,
	//Args: cobra.ExactArgs(1),
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return fmt.Errorf("missing branch argument\n")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		repoFlag, _ := cmd.Flags().GetString("repo")
		groupFlag, _ := cmd.Flags().GetString("group")
		filterFlag, _ := cmd.Flags().GetString("filter")
		kcgCmd := kcg.Command(config)

		for index, repo := range kcgCmd.List(repoFlag, groupFlag, filterFlag) {
			if path, exists := kcgCmd.Path(repo); !exists {
				fmt.Println("invalid path: '" + index + "' " + path)
				continue
			}

			if err := kcgCmd.Switch(repo, args[0]); err != nil {
				fmt.Println(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(switchCmd)
	switchCmd.Flags().String("repo", "", "repository name")
	switchCmd.Flags().String("group", "", "repository group name")
	switchCmd.Flags().String("filter", "", "repository filter")
}
