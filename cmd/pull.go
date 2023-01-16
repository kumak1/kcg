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
	"github.com/kumak1/kcg/kcg"
	"github.com/spf13/cobra"
	"sync"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "run `git pull` on each repository dir",
	Long:  `Running git pull command on each repository dir`,
	Run: func(cmd *cobra.Command, args []string) {
		groupFlag, _ := cmd.Flags().GetString("group")
		filterFlag, _ := cmd.Flags().GetString("filter")
		kcgCmd := kcg.Command(config)

		var wg sync.WaitGroup

		for index, repo := range kcgCmd.List(groupFlag, filterFlag) {
			wg.Add(1)
			repo := repo
			index := index
			go func() {
				output, err := kcgCmd.Pull(repo)

				if err == nil {
					cmd.Printf(validMessageFormat, "✔", index)
					if output != "Already up to date." {
						cmd.Println(output)
					}
				} else {
					cmd.Printf(invalidMessageFormat, "X", index)
					if output != "" {
						cmd.Println(output)
						cmd.Println(err.Error())
					} else {
						cmd.Print(err.Error())
					}
				}
				wg.Done()
			}()
		}
		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
	assignSearchFlags(pullCmd)
}
