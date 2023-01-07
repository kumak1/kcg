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
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// cloneCmd represents the clone command
var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "run `git clone` each repository",
	Long:  `Running git clone command each repository`,
	Run: func(cmd *cobra.Command, args []string) {
		groupFlag, _ := cmd.Flags().GetString("group")
		filterFlag, _ := cmd.Flags().GetString("filter")
		kcgCmd := kcg.Command(config)

		for _, repo := range kcgCmd.List(groupFlag, filterFlag) {
			if err := kcgCmd.Clone(repo); err != nil {
				fmt.Println(err)
			}
		}

		if config.Ghq {
			for index, repo := range kcgCmd.List("", "") {
				if repo.Path != "" {
					continue
				}

				fmt.Println(index)
				if path, err := kcgExec.GhqPath(repo.Repo); err == nil {
					config.Repos[index].Path = path
					fmt.Println(path)
				}
			}
			viper.Set("repos", config.Repos)
			WriteConfig("")
		}
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)
	cloneCmd.Flags().String("group", "", "repository group name")
	cloneCmd.Flags().String("filter", "", "repository filter")
}
